**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-04-22 08:48:34 UTC（2026-04-22 16:48:34 UTC+8）

**代理总数：92**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 92 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 208.87.243.199:7878 | ✓ 350ms | ✓ 1282ms | ✓ 1023ms | ✓ 841ms | ✓ 625ms | http |
| 1.231.81.166:3128 | ✓ 1367ms | ✓ 1047ms | ✓ 819ms | ✓ 892ms | ✓ 723ms | http |
| 45.76.207.177:40000 | ✓ 1362ms | 否 | ✓ 588ms | ✓ 1145ms | ✓ 812ms | http |
| 152.42.208.139:8118 | ✓ 725ms | 否 | ✓ 1181ms | ✓ 1049ms | ✓ 834ms | http |
| 46.101.95.183:8888 | ✓ 1169ms | 否 | ✓ 1229ms | ✓ 1845ms | ✓ 1277ms | http |
| 113.160.132.26:8080 | ✓ 1785ms | ✓ 1362ms | ✓ 1332ms | ✓ 1315ms | ✓ 1040ms | http |
| 115.231.181.40:8128 | ✓ 1623ms | ✓ 1424ms | 否 | ✓ 1154ms | ✓ 988ms | http |
| 188.246.224.49:7890 | ✓ 1190ms | ✓ 1805ms | ✓ 1091ms | 否 | ✓ 1775ms | http |
| 91.99.15.45:2095 | ✓ 1165ms | ✓ 1887ms | ✓ 1555ms | 否 | ✓ 1767ms | http |
| 35.225.22.61:80 | ✓ 897ms | ✓ 1696ms | ✓ 1711ms | ✓ 1287ms | ✓ 1123ms | http |
| 162.240.154.26:3128 | ✓ 902ms | 否 | 否 | ✓ 983ms | ✓ 1155ms | http |
| 14.143.222.113:57788 | ✓ 1187ms | 否 | ✓ 977ms | ✓ 1402ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1158ms | 否 | ✓ 1446ms | ✓ 1493ms | ✓ 1381ms | http |
| 155.212.188.205:8080 | ✓ 1259ms | 否 | ✓ 1767ms | 否 | ✓ 1937ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1184ms | ✓ 1215ms | ✓ 923ms | ✓ 1057ms | http |
| 161.97.184.191:8080 | ✓ 1632ms | 否 | ✓ 1125ms | ✓ 1970ms | ✓ 1353ms | http |
| 210.223.44.230:3128 | ✓ 1979ms | ✓ 897ms | ✓ 1067ms | 否 | ✓ 1566ms | http |
| 223.84.151.86:30005 | ✓ 1240ms | ✓ 1251ms | ✓ 988ms | ✓ 1205ms | ✓ 1360ms | http |
| 59.46.216.131:30001 | ✓ 988ms | ✓ 1328ms | 否 | ✓ 1613ms | ✓ 1076ms | http |
| 78.11.96.22:8888 | ✓ 1113ms | 否 | ✓ 1495ms | ✓ 1759ms | ✓ 1829ms | http |
| 85.190.99.143:443 | ✓ 1224ms | 否 | ✓ 1570ms | 否 | ✓ 1195ms | http |
| 218.153.163.156:8999 | ✓ 1976ms | 否 | ✓ 963ms | ✓ 965ms | ✓ 961ms | http |
| 34.71.229.255:3128 | ✓ 411ms | 否 | ✓ 1288ms | ✓ 1422ms | ✓ 1006ms | http |
| 43.155.86.199:8080 | ✓ 1558ms | 否 | 否 | ✓ 1276ms | ✓ 1146ms | http |
| 168.144.75.9:3128 | ✓ 1847ms | 否 | ✓ 1774ms | 否 | ✓ 1637ms | http |
| 177.93.132.244:3128 | ✓ 934ms | 否 | ✓ 1182ms | 否 | ✓ 1807ms | http |
| 8.140.104.98:3128 | ✓ 918ms | ✓ 1147ms | ✓ 952ms | ✓ 1223ms | ✓ 1023ms | http |
| 152.70.91.193:40000 | ✓ 940ms | 否 | ✓ 1030ms | ✓ 997ms | ✓ 1001ms | http |
| 213.32.85.26:3128 | ✓ 703ms | 否 | ✓ 655ms | 否 | ✓ 1564ms | http |
| 43.132.188.134:443 | ✓ 1121ms | 否 | ✓ 1434ms | 否 | ✓ 854ms | http |
| 84.47.150.126:1080 | ✓ 1969ms | 否 | ✓ 1586ms | 否 | ✓ 1979ms | http |
| 103.113.70.189:1082 | ✓ 333ms | 否 | ✓ 886ms | 否 | ✓ 937ms | http |
| 160.250.4.245:1 | ✓ 1698ms | 否 | ✓ 1296ms | ✓ 1370ms | ✓ 1000ms | http |
| 91.233.223.147:3128 | ✓ 1296ms | 否 | ✓ 977ms | 否 | ✓ 1690ms | http |
| 218.153.163.156:8993 | ✓ 1806ms | 否 | ✓ 1014ms | 否 | ✓ 1619ms | http |
| 117.236.124.166:3128 | ✓ 1866ms | 否 | ✓ 1963ms | 否 | ✓ 1853ms | http |
| 138.124.99.216:8888 | 否 | ✓ 1888ms | ✓ 1619ms | 否 | ✓ 1412ms | http |
| 5.161.196.81:8888 | ✓ 1163ms | ✓ 1538ms | ✓ 1139ms | 否 | 否 | http |
| 168.222.254.136:8888 | ✓ 1240ms | ✓ 1830ms | ✓ 1812ms | ✓ 1933ms | ✓ 1860ms | http |
| 218.153.163.156:8008 | ✓ 1872ms | 否 | ✓ 1868ms | ✓ 1853ms | ✓ 1730ms | http |
| 84.47.150.125:1080 | ✓ 836ms | 否 | ✓ 1434ms | 否 | ✓ 1634ms | http |
| 218.108.131.186:17890 | ✓ 1673ms | 否 | ✓ 804ms | ✓ 1624ms | ✓ 888ms | http |
| 37.187.109.70:10111 | ✓ 1707ms | ✓ 1781ms | ✓ 1101ms | 否 | 否 | http |
| 152.32.132.190:7890 | 否 | 否 | ✓ 927ms | ✓ 1453ms | ✓ 995ms | http |
| 112.163.160.93:3128 | ✓ 1414ms | ✓ 1473ms | ✓ 966ms | ✓ 1750ms | 否 | http |
| 45.140.147.155:1082 | ✓ 753ms | ✓ 1497ms | ✓ 1179ms | 否 | ✓ 1295ms | http |
| 45.140.147.155:1081 | ✓ 766ms | 否 | ✓ 672ms | ✓ 1771ms | 否 | http |
| 121.230.9.160:1080 | 否 | 否 | ✓ 1125ms | ✓ 1554ms | ✓ 1027ms | http |
| 147.45.186.28:3128 | 否 | 否 | ✓ 1765ms | ✓ 1636ms | ✓ 1508ms | http |
| 139.159.97.82:10900 | ✓ 1238ms | 否 | ✓ 1415ms | ✓ 1564ms | ✓ 1159ms | http |
| 168.110.52.228:3128 | ✓ 1852ms | 否 | ✓ 1520ms | ✓ 882ms | ✓ 652ms | http |
| 160.250.134.143:3128 | ✓ 922ms | 否 | ✓ 920ms | ✓ 1317ms | ✓ 966ms | http |
| 172.208.25.199:3128 | ✓ 927ms | 否 | 否 | ✓ 1559ms | ✓ 1335ms | http |
| 20.78.213.56:80 | ✓ 1456ms | 否 | ✓ 1437ms | ✓ 1347ms | ✓ 1960ms | http |
| 103.227.187.11:6090 | ✓ 1960ms | 否 | ✓ 1842ms | ✓ 1581ms | ✓ 1467ms | http |
| 167.71.196.178:80 | ✓ 755ms | 否 | ✓ 989ms | ✓ 1095ms | ✓ 852ms | http |
| 137.59.47.73:3128 | ✓ 1075ms | 否 | ✓ 1390ms | ✓ 1159ms | ✓ 1287ms | http |
| 115.178.49.111:8080 | 否 | 否 | ✓ 1753ms | ✓ 1646ms | ✓ 1466ms | http |
| 47.84.73.61:1080 | ✓ 728ms | 否 | ✓ 861ms | ✓ 1044ms | ✓ 835ms | http |
| 195.26.224.49:3128 | ✓ 1273ms | 否 | ✓ 757ms | ✓ 1612ms | ✓ 1486ms | http |
| 152.42.177.32:8888 | ✓ 732ms | 否 | ✓ 1220ms | ✓ 1256ms | ✓ 1225ms | http |
| 62.113.119.14:8080 | ✓ 736ms | ✓ 1853ms | ✓ 784ms | ✓ 1983ms | ✓ 1448ms | http |
| 103.113.70.189:1081 | ✓ 420ms | ✓ 1494ms | ✓ 265ms | 否 | ✓ 859ms | http |
| 45.140.147.82:1081 | ✓ 761ms | ✓ 1667ms | ✓ 1246ms | ✓ 1814ms | ✓ 1167ms | http |
| 185.191.236.162:3128 | ✓ 1367ms | ✓ 1815ms | 否 | 否 | ✓ 1974ms | http |
| 207.254.71.62:8088 | ✓ 1758ms | 否 | ✓ 1605ms | ✓ 1944ms | 否 | http |
| 82.114.228.67:1080 | ✓ 1394ms | 否 | 否 | ✓ 1737ms | ✓ 1249ms | http |
| 45.153.231.229:8080 | ✓ 1221ms | 否 | ✓ 1447ms | 否 | ✓ 1850ms | http |
| 89.208.106.138:10808 | ✓ 688ms | 否 | ✓ 1928ms | ✓ 1983ms | ✓ 1551ms | http |
| 20.205.16.149:3128 | ✓ 699ms | 否 | ✓ 953ms | ✓ 1182ms | ✓ 872ms | http |
| 120.132.52.172:8888 | 否 | 否 | ✓ 1588ms | ✓ 1368ms | ✓ 1148ms | http |
| 138.124.108.176:3128 | ✓ 1853ms | 否 | ✓ 952ms | ✓ 1926ms | ✓ 1216ms | http |
| 149.62.191.202:3128 | 否 | ✓ 1972ms | ✓ 740ms | 否 | ✓ 1523ms | http |
| 47.95.231.180:8084 | ✓ 1922ms | ✓ 1147ms | ✓ 936ms | ✓ 1198ms | ✓ 1030ms | http |
| 45.186.6.104:3128 | ✓ 1424ms | ✓ 1612ms | ✓ 1822ms | 否 | 否 | http |
| 20.120.225.109:3128 | 否 | 否 | ✓ 1045ms | ✓ 1549ms | ✓ 889ms | http |
| 150.249.255.91:3128 | ✓ 1592ms | 否 | ✓ 1114ms | ✓ 929ms | ✓ 710ms | http |
| 218.153.163.156:8145 | 否 | 否 | ✓ 1816ms | ✓ 1418ms | ✓ 1669ms | http |
| 175.194.173.105:3128 | 否 | ✓ 1598ms | ✓ 957ms | ✓ 1831ms | 否 | http |
| 115.248.66.131:3129 | ✓ 1608ms | 否 | 否 | ✓ 1948ms | ✓ 1885ms | http |
| 45.140.147.82:1082 | ✓ 605ms | ✓ 1418ms | 否 | ✓ 1784ms | ✓ 1268ms | http |
| 103.18.78.250:1111 | 否 | 否 | ✓ 1952ms | ✓ 1616ms | ✓ 1310ms | http |
| 178.63.155.151:8888 | ✓ 1394ms | 否 | ✓ 941ms | 否 | ✓ 1296ms | http |
| 61.52.131.172:8443 | ✓ 862ms | ✓ 1086ms | ✓ 940ms | ✓ 1113ms | ✓ 935ms | http |
| 103.184.99.194:8080 | ✓ 1773ms | 否 | 否 | ✓ 1954ms | ✓ 1406ms | http |
| 121.147.253.205:3136 | 否 | 否 | ✓ 840ms | ✓ 976ms | ✓ 1593ms | http |
| 65.108.203.37:18080 | ✓ 1355ms | 否 | 否 | ✓ 1794ms | ✓ 1533ms | http |
| 117.122.240.82:3338 | ✓ 1540ms | 否 | ✓ 954ms | ✓ 1374ms | ✓ 1057ms | http |
| 94.72.109.196:3129 | ✓ 1237ms | ✓ 1920ms | ✓ 1728ms | 否 | 否 | http |
| 103.39.51.207:8080 | ✓ 1376ms | 否 | ✓ 1756ms | ✓ 1658ms | ✓ 1871ms | http |
| 160.250.5.22:1 | ✓ 1558ms | 否 | ✓ 1091ms | ✓ 1250ms | ✓ 1069ms | http |
| 34.101.184.164:3128 | ✓ 1653ms | 否 | ✓ 1161ms | ✓ 1349ms | ✓ 1045ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
