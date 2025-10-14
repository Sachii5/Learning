import pandas as pd
import numpy as np
import matplotlib.pyplot as plt


table = pd.read_csv("annual-enterprise-survey-2024-financial-year-provisional.csv")
table.head()

x_label = table['Nama Industri']
plt.bar(x=np.arange(len(x_label)),height=table['Year'])
plt.xticks(np.arange(len(x_label)), table['Nama Industri'], rotation=90)
plt.xlabel('Nama industri')
plt.ylabel('Tahun')
plt.title('Persebaran Jumlah Penduduk Laki- Laki di Jakarta Pusat')

plt.show()
