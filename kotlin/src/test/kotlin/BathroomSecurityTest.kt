import com.willmadison.adventofcode.decodeBathRoomInstructions
import junit.framework.TestCase.assertEquals
import org.junit.Test

class BathroomSecurityTest {

    @Test
    fun testDecodeBathroomInstructions() {
        assertEquals(1985, decodeBathRoomInstructions("""ULL
                RRDDD
                LURDL
                UUUUD"""))
    }

}