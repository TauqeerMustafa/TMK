import asyncio
from typing import Dict, Any, List

class TelemetryEventProcessor:
    """
    Asynchronous event batching and metric aggregation pipeline.
    """

    def __init__(self, batch_size: int = 100):
        self.batch_size = batch_size
        self.queue: asyncio.Queue = asyncio.Queue()
        self.processed_count = 0

    async def ingest(self, event: Dict[str, Any]) -> None:
        await self.queue.put(event)

    async def process_batch(self) -> List[Dict[str, Any]]:
        batch = []
        while not self.queue.empty() and len(batch) < self.batch_size:
            item = await self.queue.get()
            batch.append(item)
            self.queue.task_done()
        self.processed_count += len(batch)
        return batch
