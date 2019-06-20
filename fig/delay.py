import numpy as np
import matplotlib.pyplot as plt

x = [2, 4, 6, 8, 10, 12, 14, 16, 18, 20]

y_greedy = [0.0358, 0.0416, 0.0415, 0.0467, 0.0491, 0.0587, 0.0678, 0.0673, 0.0729, 0.0725]
y_min    = [0.0361, 0.0478, 0.0541, 0.0612, 0.0611, 0.0620, 0.0697, 0.0762, 0.0816, 0.0817]
y_basic  = [0.0336, 0.0490, 0.0594, 0.0664, 0.0757, 0.0739, 0.0886, 0.0899, 0.0901, 0.0945]

plt.plot(x, y_greedy, label='Greedy-r',marker='o')
plt.plot(x, y_min, label='Min',marker='x')
plt.plot(x, y_basic, label='Basic',marker='s')
plt.xlabel("sta number")
plt.xticks([0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22],
	['0', '2', '4', '6', '8', '10', '12', '14', '16', '18', '20', '22'])
plt.ylabel("average interest delay (second)")
plt.legend(loc=2)
plt.show()



