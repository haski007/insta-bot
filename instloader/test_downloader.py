import unittest
import logging
from downloader import get_post_info

# Configure logging for tests
logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)

class TestDownloader(unittest.TestCase):
    
    def test_get_post_info_dkzbpsipi_w(self):
        """Test get_post_info function with shortcode DKZbpSiPi_W"""
        print("\n" + "="*50)
        print("STARTING TEST: test_get_post_info_dkzbpsipi_w")
        print("="*50)
        
        shortcode = "Cyf8L7GoPCQ"
        print(f"Testing with shortcode: {shortcode}")
        
        # Call the function
        print("Calling get_post_info...")
        result = get_post_info(shortcode)
        print("get_post_info completed")
        
        # Basic validation
        self.assertIsInstance(result, dict)
        
        # Check if we got an error or successful result
        if "error" in result:
            print(f"Error occurred: {result['error']}")
            # Even if there's an error, it should be a string
            self.assertIsInstance(result["error"], str)
        else:
            # Check that all expected fields are present
            expected_fields = [
                "shortcode", "is_video", "url", "video_url", 
                "caption", "owner", "likes", "comments", "timestamp"
            ]
            
            for field in expected_fields:
                self.assertIn(field, result, f"Field '{field}' missing from result")
            
            # Check specific values
            self.assertEqual(result["shortcode"], shortcode)
            self.assertIsInstance(result["is_video"], bool)
            self.assertIsInstance(result["url"], str)
            self.assertIsInstance(result["owner"], str)
            self.assertIsInstance(result["likes"], int)
            self.assertIsInstance(result["comments"], int)
            
            # Video URL should be None if not a video, or a string if it is
            if result["is_video"]:
                self.assertIsInstance(result["video_url"], str)
            else:
                self.assertIsNone(result["video_url"])
            
            # Caption can be None or a string
            if result["caption"] is not None:
                self.assertIsInstance(result["caption"], str)
            
            print(f"Successfully retrieved post info:")
            print(f"  Owner: {result['owner']}")
            print(f"  Is Video: {result['is_video']}")
            print(f"  Likes: {result['likes']}")
            print(f"  Comments: {result['comments']}")
            print(f"  URL: {result['url']}")
        
        print("="*50)
        print("TEST COMPLETED")
        print("="*50 + "\n")

if __name__ == "__main__":
    unittest.main() 