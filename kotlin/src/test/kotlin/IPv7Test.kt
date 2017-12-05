import com.willmadison.adventofcode.supportsTLS
import junit.framework.TestCase.assertFalse
import junit.framework.TestCase.assertTrue
import org.junit.Test

class IPv7Test {

    @Test
    fun testTLSSupport() {
        assertTrue("abba[mnop]qrst".supportsTLS())
        assertFalse("abcd[bddb]xyyx".supportsTLS())
        assertFalse("aaaa[qwer]tyui".supportsTLS())
        assertTrue("ioxxoj[asdfgh]zxcvbn".supportsTLS())
    }
}