import com.willmadison.adventofcode.decodeBathRoomInstructions
import junit.framework.Assert
import org.junit.Test

class BathroomSecurityTest {

    @Test
    fun testDecodeBathroomInstructions() {
        Assert.assertEquals(1985, decodeBathRoomInstructions("""ULL
                RRDDD
                LURDL
                UUUUD"""))
    }

}