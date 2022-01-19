
from typing import List, Tuple
import sys


class Crabs():
    def __init__(self, positions: List[int]):
        self.positions = positions

    def __calc_fuel(self, p) -> int:
        fuels = 0
        for pos in self.positions:
           #fuels += abs(pos - p) 
           n = abs(pos - p)
           fuels += int(((1 + n) * n) /2)
        return fuels
            

    def find_best_position_1(self) -> Tuple[int, int]:
        _min = min(self.positions)
        _max = max(self.positions)

        least_fuels = sys.maxsize
        least_position = _min
        for i in range(_min, _max):
            fuels = self.__calc_fuel(i)
            if fuels < least_fuels:
                least_fuels = fuels
                least_position = i
        return least_position, least_fuels

    def find_best_position_2(self) -> Tuple[int, int]:
        _min = min(self.positions)
        _max = max(self.positions)

        while _max - _min > 1:
            min_fuels = self.__calc_fuel(_min)
            max_fuels = self.__calc_fuel(_max)

            _mid = int((_min + _max)/2)
            if min_fuels < max_fuels:
                _max = _mid
            else:
                _min = _mid
        min_fuels = self.__calc_fuel(_min)
        max_fuels = self.__calc_fuel(_max)
        if min_fuels < max_fuels:
            return _min, min_fuels
        else:
            return _max, max_fuels
        

def read_positions(file_path: str) -> List[int]:
    with open(file_path) as f:
        line = f.readline()
        positions = []
        for s in line.split(","):
            positions.append(int(s))
        return positions


if __name__ == "__main__":
    #positions = [16,1,2,0,4,2,7,1,2,14]
    positions = read_positions("./crab/input.txt")
    print(positions)
    crabs = Crabs(positions)
    print(crabs.find_best_position_2())
