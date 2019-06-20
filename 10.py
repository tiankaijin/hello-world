import numpy as np
import matplotlib.pyplot as plt

#x = np.linspace(0, 10, 500)
x = [5,10,15,20,25]
y_util = [9.6, 9.75, 14.65, 23.06, 82.66]
y_mini = [9.7, 17.17, 25.21, 74.11, 102.31]
y_basic = [10.36, 18.43, 25.62, 91.91, 101.71]
y_unicast = [9.4, 15.72, 70, 93.3, 136.5]

fig, ax = plt.subplots()

# Using set_dashes() to modify dashing of an existing line
line1, = ax.plot(x, y_util, label='mda')
line2, = ax.plot(x, y_mini, label='min')
line3, = ax.plot(x, y_basic, label='basic')
line4, = ax.plot(x, y_unicast, label='unicast')
x0=20
y0=0
x1=23.06
y1=93.3
plt.plot([x0,x0,],[x1,y1],'k--',linewidth=1.5)
plt.annotate('23.06',xy=(20,23.06),xycoords='data',xytext=(+15,-15),
             textcoords='offset points',fontsize=10,
             arrowprops=dict(arrowstyle='->',connectionstyle='arc3,rad=.2'))
plt.annotate('(74.11, %%%.1f)' % ((74.11-23.06)*100/74.11),xy=(20,74.11),xycoords='data',xytext=(+15,-15),
             textcoords='offset points',fontsize=10,
             arrowprops=dict(arrowstyle='->',connectionstyle='arc3,rad=.2'))
plt.annotate('(91.91, %%%.1f)' % ((91.91-23.06)*100/91.91),xy=(20,91.91),xycoords='data',xytext=(+10,+20),
             textcoords='offset points',fontsize=10,
             arrowprops=dict(arrowstyle='->',connectionstyle='arc3,rad=.2'))
plt.annotate('(93.3, %%%.1f)' % ((93.3-23.06)*100/93.3),xy=(20,93.3),xycoords='data',xytext=(-30,+30),
             textcoords='offset points',fontsize=10,
             arrowprops=dict(arrowstyle='->',connectionstyle='arc3,rad=.2'))
plt.xlabel("sta number")
plt.xticks([5,10,15,20,25],
  ['5', '10', '15', '20', '25'])
plt.ylabel("last finish time(second)")
#line1.set_dashes([2, 2, 10, 2])  # 2pt line, 2pt break, 10pt line, 2pt break

# Using plot(..., dashes=...) to set the dashing when creating a line
#line2, = ax.plot(x, y - 0.2, dashes=[6, 2], label='Using the dashes parameter')

ax.legend()
plt.show()