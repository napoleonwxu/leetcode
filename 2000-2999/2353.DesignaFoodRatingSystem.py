from sortedcontainers import SortedList

class FoodRatings:

    def __init__(self, foods: List[str], cuisines: List[str], ratings: List[int]):
        self.food_info = {}
        self.cuisine_to_foods = defaultdict(SortedList)
        for food, cuisine, rate in zip(foods, cuisines, ratings):
            self.food_info[food] = (cuisine, rate)
            self.cuisine_to_foods[cuisine].add((-rate, food))
        

    def changeRating(self, food: str, newRating: int) -> None:
        cuisine, rate = self.food_info[food]
        self.food_info[food] = (cuisine, newRating)
        self.cuisine_to_foods[cuisine].remove((-rate, food))
        self.cuisine_to_foods[cuisine].add((-newRating, food))

    def highestRated(self, cuisine: str) -> str:
        return self.cuisine_to_foods[cuisine][0][1]


# Your FoodRatings object will be instantiated and called as such:
# obj = FoodRatings(foods, cuisines, ratings)
# obj.changeRating(food,newRating)
# param_2 = obj.highestRated(cuisine)