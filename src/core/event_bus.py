# TMK Event Bus
import collections

class EventBus:
    def __init__(self):
        self._subscribers = collections.defaultdict(list)

    def subscribe(self, topic: str, handler):
        self._subscribers[topic].append(handler)

    def publish(self, topic: str, data=None):
        for handler in self._subscribers.get(topic, []):
            handler(data)
