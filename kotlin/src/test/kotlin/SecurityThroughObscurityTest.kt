import com.willmadison.adventofcode.parseRoom
import junit.framework.TestCase.assertEquals
import junit.framework.TestCase.assertTrue
import org.junit.Test

class SecurityThroughObscurityTest {

    @Test
    fun testParseEncryptedName() {
        val (room, givenChecksum) = "aaaaa-bbb-z-y-x-123[abxyz]".parseRoom()

        assertEquals(room.sectorId, 123)
        assertEquals(room.checksum, "abxyz")
        assertEquals(room.checksum, givenChecksum)
        assertTrue(room.isReal(givenChecksum))
    }
}

