##### Problem Statement: Two Crystal Balls

You are given two identical crystal balls and a building with `n` floors. You need to determine the highest floor from which you can drop a crystal ball without it breaking.

**Conditions:**

1. If you drop a crystal ball from a floor `f` and it doesn't break, then it will not break if dropped from any floor below `f`.
2. If you drop a crystal ball from a floor `f` and it breaks, then it will break if dropped from any floor above `f`.
3. You have only two crystal balls. If both crystal balls break, you cannot perform any more tests.

**Objective:**

Determine the highest safe floor (the highest floor from which you can drop a crystal ball without it breaking) with the minimum number of drops in the worst-case scenario.

**Constraints:**

- You must minimize the number of drops in the worst-case scenario.
- You are allowed to break one crystal ball in the process of determining the highest safe floor.

**Output:**

Return the highest floor number `f` from which you can safely drop a crystal ball without it breaking.

**Example:**

- If the building has 100 floors and the crystal ball does not break when dropped from the 10th floor but breaks when dropped from the 11th floor, the highest safe floor is 10.
