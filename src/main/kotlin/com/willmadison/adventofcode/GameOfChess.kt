package com.willmadison.adventofcode

import java.security.MessageDigest
import java.math.BigInteger

enum class Strategy {
    BASIC, MISDIRECTION
}

fun derivePassword(doorId: String, strategy: Strategy = Strategy.BASIC): String {
    val characters = mutableListOf<Char>()
    val charactersByIndex = mutableMapOf<Char, Char>()

    var i = 0
    var finished: Boolean

    do {
        val m = MessageDigest.getInstance("MD5")
        m.update("$doorId$i".toByteArray())
        val digest = m.digest()
        var hash = BigInteger(1, digest).toString(16)

        while (hash.length < 32) {
            hash = "0$hash"
        }

        if (hash.startsWith("00000")) {
            when (strategy) {
                Strategy.BASIC -> characters.add(hash[5])
                Strategy.MISDIRECTION -> {
                    val position = hash[5]

                    if (position in '0'..'7' && !charactersByIndex.containsKey(position)) {
                        charactersByIndex[position] = hash[6]
                    }
                }
            }
        }

        i++
        finished = characters.size == 8 || charactersByIndex.size == 8
    } while (!finished)

    return when {
        characters.size == 8 -> characters.joinToString("", transform = Char::toString)
        else -> {
            val buffer = StringBuffer()

            for (index in '0'..'7') {
                buffer.append(charactersByIndex[index])
            }

            buffer.toString()
        }
    }
}

fun main(args: Array<String>) {
    println("The password for my door id is ${derivePassword("reyedfim", Strategy.MISDIRECTION)}")
}