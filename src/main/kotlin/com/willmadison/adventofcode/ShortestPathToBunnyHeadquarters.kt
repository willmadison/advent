package com.willmadison.adventofcode

import org.slf4j.LoggerFactory

enum class Direction {
    NORTH, SOUTH, EAST, WEST
}

data class Coordinate(var x: Int = 0, var y: Int = 0)

class Vector(var direction: Char, var magnitude: Int)

class Person(val currentPosition: Coordinate = Coordinate(), var currentDirection: Direction = Direction.NORTH) {

    fun move(vector: Vector) {
        turn(vector)
        walk(vector)
    }

    private fun turn(vector: Vector) {
        when (vector.direction) {
            'R' -> {
                when (currentDirection) {
                    Direction.NORTH -> currentDirection = Direction.EAST
                    Direction.SOUTH -> currentDirection = Direction.WEST
                    Direction.EAST -> currentDirection = Direction.SOUTH
                    Direction.WEST -> currentDirection = Direction.NORTH
                }
            }
            'L' -> {
                when (currentDirection) {
                    Direction.NORTH -> currentDirection = Direction.WEST
                    Direction.SOUTH -> currentDirection = Direction.EAST
                    Direction.EAST -> currentDirection = Direction.NORTH
                    Direction.WEST -> currentDirection = Direction.SOUTH
                }
            }
        }
    }

    private fun walk(vector: Vector) {
        when (currentDirection) {
            Direction.NORTH -> currentPosition.y += vector.magnitude
            Direction.SOUTH -> currentPosition.y -= vector.magnitude
            Direction.EAST -> currentPosition.x += vector.magnitude
            Direction.WEST -> currentPosition.x -= vector.magnitude
        }
    }

    fun distanceTravelled() = currentPosition.x.abs() + currentPosition.y.abs()
}

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

    val person = Person()

    steps.forEach { step ->
        val vector = step.parseVector()
        person.move(vector)
    }

    return person.distanceTravelled()
}

fun main(args: Array<String>) {
    val logger = LoggerFactory.getLogger("com.willmadison.adventofcode.shortestPath")
    val path = "L4, R2, R4, L5, L3, L1, R4, R5, R1, R3, L3, L2, L2, R5, R1, L1, L2, R2, R2, L5, R5, R5, L2, R1, R2, L2, L4, L1, R5, R2, R1, R1, L2, L3, R2, L5, L186, L5, L3, R3, L5, R4, R2, L5, R1, R4, L1, L3, R3, R1, L1, R4, R2, L1, L4, R5, L1, R50, L4, R3, R78, R4, R2, L4, R3, L4, R4, L1, R5, L4, R1, L2, R3, L2, R5, R5, L4, L1, L2, R185, L5, R2, R1, L3, R4, L5, R2, R4, L3, R4, L2, L5, R1, R2, L2, L1, L2, R2, L2, R1, L5, L3, L4, L3, L4, L2, L5, L5, R2, L3, L4, R4, R4, R5, L4, L2, R4, L5, R3, R1, L1, R3, L2, R2, R1, R5, L4, R5, L3, R2, R3, R1, R4, L4, R1, R3, L5, L1, L3, R2, R1, R4, L4, R3, L3, R3, R2, L3, L3, R4, L2, R4, L3, L4, R5, R1, L1, R5, R3, R1, R3, R4, L1, R4, R3, R1, L5, L5, L4, R4, R3, L2, R1, R5, L3, R4, R5, L4, L5, R2"

    logger.info("The shortest path to Bunny Headquarters is: ${shortestPath(path)}")
}