import unittest
from rotate import rotate


class TestRotate(unittest.TestCase):

    def test_rotate(self):
        assert rotate("hello", 2) == "llohe"
        assert rotate("hello", -2) == "lohel"
        assert rotate("", 3) == ""


if __name__ == '__main__':
    unittest.main()