package com.willmadison.adventofcode

import java.io.File

fun List<String>.findHyperNetSequences(source: String): Set<String> {
    return this.filter { source.contains("[$it]") }.toSet()
}

fun String.isABBASequence(): Boolean {
    var (i, j) = 0 to this.length - 1

    while (i < j) {
        if (this[i] != this[j]) {
            return false
        }

        i++
        j--
    }

    return this[0] != this[1]
}

fun String.supportsTLS(): Boolean {
    val parts = this.split("[", "]")
    val hypernetSequences = parts.findHyperNetSequences(this)

    for (sequence in hypernetSequences) {
        for (i in 0..sequence.length - 4) {
            val s = sequence.substring(i..i + 3)

            if (s.isABBASequence()) {
                return false
            }
        }
    }

    val nonHypernetSequences = parts.filter { !hypernetSequences.contains(it) }

    for (sequence in nonHypernetSequences) {
        for (i in 0..sequence.length - 4) {
            val s = sequence.substring(i..i + 3)

            if (s.isABBASequence()) {
                return true
            }
        }
    }

    return false
}

fun main(args: Array<String>) {
    val ipAddresses = File("day7_input").readLines().filter(String::supportsTLS).count()
    println("There are $ipAddresses IPs supporting TLS.")
}