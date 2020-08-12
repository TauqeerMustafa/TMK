import unittest
from src.core.cache import LRUCache
from src.security.crypto import generate_signature, verify_signature

class TestTMK(unittest.TestCase):
    def test_cache(self):
        c = LRUCache(2)
        c.put("a", 1)
        c.put("b", 2)
        self.assertEqual(c.get("a"), 1)
        c.put("c", 3)
        self.assertIsNone(c.get("b"))

    def test_security(self):
        sig = generate_signature("secret", "hello")
        self.assertTrue(verify_signature("secret", "hello", sig))
        self.assertFalse(verify_signature("wrong", "hello", sig))

if __name__ == "__main__":
    unittest.main()
