import unittest
import sys
import os

# Ensure src is in python path
sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

from src.security.hmac_auth import TokenSecurityManager

class TestTokenSecurity(unittest.TestCase):
    def setUp(self):
        self.sec = TokenSecurityManager('test_secret_key_minimum_length_123')

    def test_valid_signature(self):
        sig, ts = self.sec.sign_payload('payload_test')
        self.assertTrue(self.sec.verify_payload('payload_test', sig, ts))

    def test_tampered_payload(self):
        sig, ts = self.sec.sign_payload('payload_test')
        self.assertFalse(self.sec.verify_payload('tampered_payload', sig, ts))

    def test_api_key_generation(self):
        key = self.sec.generate_api_key()
        self.assertTrue(key.startswith('tmk_live_'))
        self.assertEqual(len(key), len('tmk_live_') + 48)

if __name__ == '__main__':
    unittest.main()
