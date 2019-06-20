import numpy as np
import matplotlib.pyplot as plt

x = [2, 4, 6, 8, 10, 12, 14, 16, 18, 20]
y_greedy = [0.9, 6.8, 6.9, 9.4, 10.2, 12.1, 13.2, 14.1, 16.8, 20.1]
y_min    = [0.08, 3.5, 6.6, 6.5, 8.1, 9.0, 9.3, 10.4, 15.0, 17.1]
y_basic  = [0.1, 4.7, 4.5, 5.7, 6.6, 6.8, 7.5, 8.2, 12.1, 15.8]


plt.plot(x, y_greedy, label='Greedy-r',marker='o')
plt.plot(x, y_min, label='Min',marker='x')
plt.plot(x, y_basic, label='Basic',marker='s')
plt.xlabel("sta number")
plt.xticks([0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22],
	['0', '2', '4', '6', '8', '10', '12', '14', '16', '18', '20', '22'])
plt.ylabel("cache hit ratio (%)")
plt.legend(loc=2)
plt.show()



