import unittest
import time
import sys
import os

# Ensure src is in python path
sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

from src.cache.lru_ttl import ThreadSafeTTLCache

class TestTTLCache(unittest.TestCase):
    def test_cache_put_get(self):
        cache = ThreadSafeTTLCache(max_size=3, default_ttl=10.0)
        cache.put('a', 1)
        cache.put('b', 2)
        self.assertEqual(cache.get('a'), 1)
        self.assertEqual(cache.get('b'), 2)

    def test_lru_eviction(self):
        cache = ThreadSafeTTLCache(max_size=2, default_ttl=10.0)
        cache.put('a', 1)
        cache.put('b', 2)
        cache.put('c', 3)
        self.assertIsNone(cache.get('a'))
        self.assertEqual(cache.get('b'), 2)
        self.assertEqual(cache.get('c'), 3)

    def test_ttl_expiration(self):
        cache = ThreadSafeTTLCache(max_size=5, default_ttl=0.05)
        cache.put('temp', 'expiring_data')
        self.assertEqual(cache.get('temp'), 'expiring_data')
        time.sleep(0.06)
        self.assertIsNone(cache.get('temp'))

if __name__ == '__main__':
    unittest.main()
