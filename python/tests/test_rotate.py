import unittest

import python.rotate as twist


class TestRotate(unittest.TestCase):

    def test_rotate(self):
        assert twist.rotate("hello", 2) == "llohe"
        assert twist.rotate("hello", -2) == "lohel"
        assert twist.rotate("", 3) == ""


if __name__ == '__main__':
    unittest.main()
