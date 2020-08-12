import hashlib
import hmac
import time
import secrets
from typing import Tuple

class TokenSecurityManager:
    """
    Cryptographic Security Token and HMAC Webhook Signature Validator.
    Protects API ingress with constant-time equality validation.
    """

    def __init__(self, master_secret: str):
        if not master_secret or len(master_secret) < 16:
            raise ValueError("Master secret must be at least 16 characters long.")
        self.secret = master_secret.encode("utf-8")

    def sign_payload(self, payload: str, timestamp: int = None) -> Tuple[str, int]:
        if timestamp is None:
            timestamp = int(time.time())
        token_base = f"{timestamp}:{payload}".encode("utf-8")
        signature = hmac.new(self.secret, token_base, hashlib.sha256).hexdigest()
        return signature, timestamp

    def verify_payload(self, payload: str, signature: str, timestamp: int, max_age_seconds: int = 300) -> bool:
        current_time = int(time.time())
        if abs(current_time - timestamp) > max_age_seconds:
            return False
        expected_sig, _ = self.sign_payload(payload, timestamp)
        return hmac.compare_digest(expected_sig, signature)

    @staticmethod
    def generate_api_key(prefix: str = "tmk_live_") -> str:
        entropy = secrets.token_hex(24)
        return f"{prefix}{entropy}"
