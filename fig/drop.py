import numpy as np
import matplotlib.pyplot as plt

x = [2, 4, 6, 8, 10, 12, 14, 16, 18, 20]
y_greedy = [0.97, 1.4, 1.9, 2.6, 2.6, 3.9, 4.2, 6.7, 6.7, 7.2]
y_min    = [1.5, 3.6, 3.8, 4.0, 4.6, 6.8, 8.8, 11.2, 12.4, 13.9]
y_basic  = [1.1, 2.5, 2.0, 2.5, 3.3, 4.1, 7.4, 9.5, 11.2, 11.4]


plt.plot(x, y_greedy, label='Greedy-r',marker='o')
plt.plot(x, y_min, label='Min',marker='x')
plt.plot(x, y_basic, label='Basic',marker='s')
plt.xlabel("sta number")
plt.xticks([0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22],
	['0', '2', '4', '6', '8', '10', '12', '14', '16', '18', '20', '22'])
plt.ylabel("loss ratio (%)")
plt.legend(loc=2)
plt.show()




