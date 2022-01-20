
from typing import List, Tuple


class Crabs():
    def __init__(self, positions: List[int]):
        self.positions = positions

    def __calc_fuel(self, p) -> int:
        fuels = 0
        for pos in self.positions:
            #fuels += abs(pos - p)
            n = abs(pos - p)
            fuels += (1 + n) * n // 2
        return fuels

    def find_best_position(self) -> Tuple[int, int]:
        _min, _max = min(self.positions), max(self.positions)
        while _max - _min > 1:
            min_fuels, max_fuels = self.__calc_fuel(
                _min), self.__calc_fuel(_max)

            _mid = (_min + _max)//2
            if min_fuels < max_fuels:
                _max = _mid
            else:
                _min = _mid
        min_fuels, max_fuels = self.__calc_fuel(_min), self.__calc_fuel(_max)
        return (_min, min_fuels) if min_fuels < max_fuels else (_max, max_fuels)


def read_positions(file_path: str) -> List[int]:
    with open(file_path) as f:
        positions = []
        for s in f.readline().split(","):
            positions.append(int(s))
        return positions


if __name__ == "__main__":
    #positions = [16,1,2,0,4,2,7,1,2,14]
    positions = read_positions("./crab/input.txt")
    crabs = Crabs(positions)
    print(crabs.find_best_position())
