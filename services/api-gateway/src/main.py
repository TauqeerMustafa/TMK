import time
from src.security.hmac_auth import TokenSecurityManager
from src.cache.lru_ttl import ThreadSafeTTLCache
from src.pipeline.event_processor import TelemetryEventProcessor

def run_app():
    print("=======================================================")
    print("⚡ TMK API Gateway & Security Ingress (Python 3.11+)")
    print("=======================================================")
    
    sec = TokenSecurityManager("production_secret_key_tmk_00192847")
    cache = ThreadSafeTTLCache(max_size=500, default_ttl=30.0)
    
    # Store initial health
    cache.put("health", {"status": "ONLINE", "uptime": time.time()})
    sig, ts = sec.sign_payload("health_check_payload")
    valid = sec.verify_payload("health_check_payload", sig, ts)
    
    print(f"✅ Cache initialized: health -> {cache.get('health')}")
    print(f"✅ Security verification: HMAC SHA256 valid = {valid}")
    print("🚀 API Gateway initialized successfully.")

if __name__ == "__main__":
    run_app()
