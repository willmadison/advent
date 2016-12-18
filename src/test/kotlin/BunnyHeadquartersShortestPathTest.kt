import com.willmadison.adventofcode.parseVector
import com.willmadison.adventofcode.shortestPath
import junit.framework.TestCase.assertEquals
import org.junit.Test

class BunnyHeadquartersShortestPathTest {

    @Test
    fun parseVector() {
        val vector = "R45".parseVector()
        assertEquals('R', vector.direction)
        assertEquals(45, vector.magnitude)
    }

    @Test
    fun shortestPathGivenTestCases() {
        assertEquals(5, shortestPath("R2, L3"))
        assertEquals(2, shortestPath("R2, R2, R2"))
        assertEquals(12, shortestPath("R5, L5, R5, R3"))
    }

}