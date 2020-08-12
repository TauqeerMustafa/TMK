# TMK DevOps System Monitor
import os
import time

def get_system_health():
    return {
        "status": "HEALTHY",
        "timestamp": time.time(),
        "pid": os.getpid()
    }
