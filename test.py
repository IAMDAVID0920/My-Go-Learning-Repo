import pandas as pd
#  Q1
num_stem_cell = [116, 117, 120, 1, 52, 79, 109, 27, 85, 51, 78, 55, 26, 39, 107]
photo_act = [60, 67, 64, 8, 13, 63, 63, 2, 46, 27, 43, 24, 10, 28, 56]
d = {"Number of Stem Cell": num_stem_cell,"Photoreceptor Activity":photo_act}
print(d)
df = pd.DataFrame(d)
print(df.head(5))
# Q2
import math
import numpy as np
import scipy
# Here the point is to get the regression coefficient 
slope, intercept, pearson_r, p_value, std_err = scipy.stats.linregress(num_stem_cell, photo_act)
print(f"coefficient of determination: {pearson_r}")
# r is regression coefficient
print(f'r = {round(pearson_r, 2)}')
# standard deviation of X
s_stem_x = np.std(num_stem_cell)
print(f'stem cell std = {round(s_stem_x, 2)}')
# standard deviation of X
s_activ_y = np.std(photo_act)
print(f'photoreceptor activity std = {round(s_activ_y, 2)}')
b = pearson_r * (s_stem_x / s_activ_y)
print(f'b = {round(b, 2)}')
# Q3 
mean_stem_x = sum(num_stem_cell) / len(num_stem_cell)
print(f'mean for stem cells = {round(mean_stem_x, 2)}')
mean_activ_y = sum(photo_act) / len(photo_act)
print(f'mean for photoreceptor activity = {round(mean_activ_y, 2)}')
print(f'a {round(intercept, 2)}')
# Q4
from sklearn.linear_model import LinearRegression
from sklearn.model_selection import train_test_split  
y = df['Photoreceptor Activity'].values.reshape(-1, 1)
X = df['Number of Stem Cell'].values.reshape(-1, 1)
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size = 1/3, random_state = 0)
model = LinearRegression()
model.fit(X_train,y_train) 
x_pred = model.predict([[70]])
x = x_pred[0][0]
print(f'Number of stem cells needed is: {round(x, 2)}')