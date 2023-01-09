import pickle

import numpy as np
from sklearn.ensemble import RandomForestRegressor

class Model:
    def __init__(self):
        self.model = RandomForestRegressor()

    def train(self, X, y):
        self.model.fit(X, y)

    def predict(self, X):
        return self.model.predict(X)

    def save(self, filepath):
        with open(filepath, "wb") as f:
            pickle.dump(self.model, f)

    def load(self, filepath):
        with open(filepath, "rb") as f:
            self.model = pickle.load(f)
