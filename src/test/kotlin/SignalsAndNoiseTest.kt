import com.willmadison.adventofcode.Cipher
import com.willmadison.adventofcode.decipherMessage
import junit.framework.TestCase.assertEquals
import org.junit.Test

class SignalsAndNoiseTest {

    @Test
    fun testDecipherMessage() {
        assertEquals("easter", decipherMessage("""eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar"""))
    }

    @Test
    fun testDecipherMessageWithModRepetition() {
        assertEquals("advent", decipherMessage("""eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar""", Cipher.MODIFIED_REPETITION))
    }

}