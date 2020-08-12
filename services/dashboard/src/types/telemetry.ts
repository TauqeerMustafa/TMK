export type HealthStatus = 'HEALTHY' | 'DEGRADED' | 'CRITICAL' | 'STANDBY';

export interface SystemMetrics {
  readonly cpuUtilization: number;
  readonly memoryUsageMb: number;
  readonly activeGoroutines: number;
  readonly throughputRps: number;
  readonly p99LatencyMs: number;
  readonly timestamp: number;
}

export interface IngressEvent {
  readonly id: string;
  readonly topic: string;
  readonly payload: Record<string, unknown>;
  readonly signature: string;
  readonly timestamp: number;
}

export interface ServiceNode {
  readonly id: string;
  readonly name: string;
  readonly runtime: 'Go' | 'Python' | 'TypeScript';
  readonly status: HealthStatus;
  readonly uptimeSeconds: number;
}
