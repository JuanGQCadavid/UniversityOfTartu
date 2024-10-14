import pandas as pd
import matplotlib.pyplot as plt

df = pd.read_csv('Resume_fib_ran.csv', delimiter=';')

fig, ax = plt.subplots()
scatter = ax.scatter(df['K'], df['DataSize'], c=df['TimeTakenSeconds'], cmap='viridis', alpha=0.75)
cbar = plt.colorbar(scatter)
cbar.set_label('Time Taken (seconds)')
ax.set_xlabel('K')
ax.set_ylabel('Data Size')
ax.set_title('K vs Data Size vs Time Taken')

plt.xscale('log')
plt.yscale('log')
plt.show()