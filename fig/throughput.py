import numpy as np
import matplotlib.pyplot as plt

x = [2, 4, 6, 8, 10, 12, 14, 16, 18, 20]
y_greedy = [11.2, 17.0, 19.0, 21.1, 22.5, 22.7, 22.4, 21.6, 21.0, 20.7]
y_min    = [11.4, 14.5, 16.3, 18.8, 19.6, 19.7, 21.1, 22.0, 21.3, 21.3]
y_basic  = [11.0, 14.0, 14.0, 14.3, 16.9, 17.2, 18.1, 19.4, 19.8, 20.1]

#fig, ax = plt.subplots()
plt.plot(x, y_greedy, label='Greedy-r',marker='o')
plt.plot(x, y_min, label='Min',marker='x')
plt.plot(x, y_basic, label='Basic',marker='s')

plt.xlabel("sta number")
plt.xticks([0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22],
	['0', '2', '4', '6', '8', '10', '12', '14', '16', '18', '20', '22'])
plt.ylabel("throughput (Mbit/s)")
plt.legend(loc=2)
plt.show()

