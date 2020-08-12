import { SystemMetrics } from '../types/telemetry';

export class LiveTelemetryStream {
  private listeners: Array<(metrics: SystemMetrics) => void> = [];
  private intervalId: ReturnType<typeof setInterval> | null = null;

  public subscribe(callback: (metrics: SystemMetrics) => void): () => void {
    this.listeners.push(callback);
    return () => {
      this.listeners = this.listeners.filter(cb => cb !== callback);
    };
  }

  public startStreaming(pollIntervalMs: number = 1000): void {
    if (this.intervalId) return;
    this.intervalId = setInterval(() => {
      const mockMetrics: SystemMetrics = {
        cpuUtilization: +(10 + Math.random() * 8).toFixed(1),
        memoryUsageMb: +(240 + Math.random() * 15).toFixed(1),
        activeGoroutines: Math.floor(60 + Math.random() * 10),
        throughputRps: Math.floor(1800 + Math.random() * 200),
        p99LatencyMs: +(1.5 + Math.random() * 0.5).toFixed(2),
        timestamp: Date.now()
      };
      this.listeners.forEach(fn => fn(mockMetrics));
    }, pollIntervalMs);
  }

  public stopStreaming(): void {
    if (this.intervalId) {
      clearInterval(this.intervalId);
      this.intervalId = null;
    }
  }
}
