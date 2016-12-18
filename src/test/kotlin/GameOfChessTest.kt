import com.willmadison.adventofcode.Strategy
import com.willmadison.adventofcode.derivePassword
import junit.framework.TestCase.assertEquals
import org.junit.Test

class GameOfChessTest {

    @Test
    fun testDerivePassword() = assertEquals("18f47a30", derivePassword("abc", Strategy.BASIC))

    @Test
    fun testDerivePasswordUsingMisdirection() = assertEquals("05ace8e3", derivePassword("abc", Strategy.MISDIRECTION))

}