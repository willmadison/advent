package com.willmadison.adventofcode

import org.slf4j.LoggerFactory

enum class Direction {
    NORTH, SOUTH, EAST, WEST
}

class Coordinate(var x: Int, var y: Int)

class Vector(var direction: Char, var magnitude: Int)

fun Int.abs(): Int {
    return if (this > 0) {
        this
    } else {
        -1 * this
    }
}

fun String.parseVector(): Vector {
    val direction = this[0]
    val magnitude = this.slice(1..this.length - 1)

    return Vector(direction, magnitude.toInt())
}

fun shortestPath(path: String): Int {
    val steps = path.split(", ")

    val currentLocation = Coordinate(0, 0)
    var currentDirection = Direction.NORTH

    steps.forEach { step ->
        val vector = step.parseVector()

        when (currentDirection) {
            Direction.NORTH -> {
                when (vector.direction) {
                    'R' -> {
                        currentDirection = Direction.EAST
                        currentLocation.x += vector.magnitude
                    }
                    'L' -> {
                        currentDirection = Direction.WEST
                        currentLocation.x -= vector.magnitude
                    }
                }
            }
            Direction.SOUTH -> {
                when (vector.direction) {
                    'R' -> {
                        currentDirection = Direction.WEST
                        currentLocation.x -= vector.magnitude
                    }
                    'L' -> {
                        currentDirection = Direction.EAST
                        currentLocation.x += vector.magnitude
                    }
                }
            }
            Direction.EAST -> {
                when (vector.direction) {
                    'R' -> {
                        currentDirection = Direction.SOUTH
                        currentLocation.y -= vector.magnitude
                    }
                    'L' -> {
                        currentDirection = Direction.NORTH
                        currentLocation.y += vector.magnitude
                    }
                }
            }
            Direction.WEST -> {
                when (vector.direction) {
                    'R' -> {
                        currentDirection = Direction.NORTH
                        currentLocation.y += vector.magnitude
                    }
                    'L' -> {
                        currentDirection = Direction.SOUTH
                        currentLocation.y -= vector.magnitude
                    }
                }
            }
        }
    }

    return currentLocation.x.abs() + currentLocation.y.abs()
}

fun main(args: Array<String>) {
    val logger = LoggerFactory.getLogger("com.willmadison.adventofcode.shortestPath")
    val path = "L4, R2, R4, L5, L3, L1, R4, R5, R1, R3, L3, L2, L2, R5, R1, L1, L2, R2, R2, L5, R5, R5, L2, R1, R2, L2, L4, L1, R5, R2, R1, R1, L2, L3, R2, L5, L186, L5, L3, R3, L5, R4, R2, L5, R1, R4, L1, L3, R3, R1, L1, R4, R2, L1, L4, R5, L1, R50, L4, R3, R78, R4, R2, L4, R3, L4, R4, L1, R5, L4, R1, L2, R3, L2, R5, R5, L4, L1, L2, R185, L5, R2, R1, L3, R4, L5, R2, R4, L3, R4, L2, L5, R1, R2, L2, L1, L2, R2, L2, R1, L5, L3, L4, L3, L4, L2, L5, L5, R2, L3, L4, R4, R4, R5, L4, L2, R4, L5, R3, R1, L1, R3, L2, R2, R1, R5, L4, R5, L3, R2, R3, R1, R4, L4, R1, R3, L5, L1, L3, R2, R1, R4, L4, R3, L3, R3, R2, L3, L3, R4, L2, R4, L3, L4, R5, R1, L1, R5, R3, R1, R3, R4, L1, R4, R3, R1, L5, L5, L4, R4, R3, L2, R1, R5, L3, R4, R5, L4, L5, R2"

    logger.info("The shortest path to Bunny Headquarters is: ${shortestPath(path)}")
}