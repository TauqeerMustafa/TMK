# TMK Security Token Validator
import hashlib
import hmac

def generate_signature(secret: str, payload: str) -> str:
    return hmac.new(secret.encode(), payload.encode(), hashlib.sha256).hexdigest()

def verify_signature(secret: str, payload: str, signature: str) -> bool:
    expected = generate_signature(secret, payload)
    return hmac.compare_digest(expected, signature)
