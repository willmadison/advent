import com.willmadison.adventofcode.determinePossibleTriangles
import com.willmadison.adventofcode.parseTriple
import junit.framework.TestCase.assertEquals
import org.junit.Test

class ThreeSidedSquaresTest {
    @Test
    fun testThreeSidedSquares() {
        assertEquals(1, determinePossibleTriangles("""5 10 25
3 4 5"""))
    }

    @Test
    fun parseTripleTest() {
        assertEquals(Triple(1, 1, 1), "1 1 1".parseTriple())
    }
}