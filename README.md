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

最后更新：2026-03-08 20:28:43 UTC（2026-03-09 04:28:43 UTC+8）

**代理总数：89**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 88 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 89 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 446ms | ✓ 1956ms | ✓ 901ms | 否 | ✓ 972ms | http |
| 14.56.107.244:3128 | ✓ 1762ms | ✓ 1231ms | ✓ 1029ms | 否 | ✓ 976ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1669ms | ✓ 1319ms | ✓ 1325ms | ✓ 1518ms | http |
| 113.177.131.2:3128 | ✓ 1601ms | ✓ 1686ms | ✓ 1005ms | ✓ 1156ms | ✓ 944ms | http |
| 202.155.12.161:443 | ✓ 1565ms | ✓ 1300ms | ✓ 1275ms | ✓ 1156ms | ✓ 903ms | http |
| 101.43.255.96:80 | ✓ 1042ms | ✓ 1318ms | ✓ 1029ms | ✓ 1206ms | ✓ 1057ms | http |
| 81.70.169.194:80 | ✓ 961ms | ✓ 1373ms | ✓ 1056ms | ✓ 1538ms | ✓ 1054ms | http |
| 185.115.74.185:8080 | ✓ 942ms | ✓ 1847ms | ✓ 1789ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 1009ms | ✓ 1228ms | ✓ 951ms | ✓ 1450ms | 否 | http |
| 152.42.213.210:8080 | 否 | 否 | ✓ 1296ms | ✓ 1073ms | ✓ 838ms | http |
| 120.92.212.16:7890 | ✓ 1221ms | ✓ 1215ms | 否 | ✓ 1252ms | ✓ 971ms | http |
| 190.9.109.207:999 | ✓ 1380ms | ✓ 1650ms | ✓ 1309ms | ✓ 1563ms | ✓ 1228ms | http |
| 190.9.109.198:999 | ✓ 1380ms | ✓ 1562ms | ✓ 1397ms | ✓ 1603ms | 否 | http |
| 194.213.18.200:443 | ✓ 857ms | ✓ 1114ms | 否 | 否 | ✓ 1202ms | http |
| 120.240.35.160:22222 | ✓ 949ms | 否 | ✓ 988ms | ✓ 1129ms | 否 | http |
| 165.227.5.10:8888 | ✓ 586ms | 否 | ✓ 1036ms | 否 | ✓ 1912ms | http |
| 222.184.48.241:22222 | ✓ 868ms | ✓ 1080ms | 否 | ✓ 1256ms | 否 | http |
| 114.231.73.92:1080 | ✓ 1309ms | ✓ 1824ms | ✓ 1659ms | ✓ 1629ms | ✓ 1277ms | http |
| 35.225.22.61:80 | ✓ 684ms | 否 | ✓ 285ms | 否 | ✓ 1137ms | http |
| 211.171.114.154:3128 | ✓ 884ms | ✓ 1356ms | 否 | 否 | ✓ 1574ms | http |
| 117.159.239.51:22222 | ✓ 1094ms | ✓ 1251ms | ✓ 949ms | ✓ 1082ms | ✓ 846ms | http |
| 125.128.12.14:3128 | ✓ 1593ms | ✓ 1521ms | ✓ 1228ms | ✓ 1012ms | ✓ 1265ms | http |
| 61.72.110.54:3128 | ✓ 657ms | ✓ 1592ms | ✓ 1277ms | ✓ 1512ms | ✓ 833ms | http |
| 152.42.213.210:80 | ✓ 1716ms | 否 | ✓ 1613ms | ✓ 1754ms | ✓ 1581ms | http |
| 116.80.82.228:3172 | 否 | 否 | ✓ 1547ms | ✓ 1869ms | ✓ 1700ms | http |
| 116.80.82.232:3172 | 否 | 否 | ✓ 1546ms | ✓ 1892ms | ✓ 1691ms | http |
| 14.103.139.54:10808 | ✓ 870ms | ✓ 1081ms | ✓ 844ms | ✓ 1162ms | ✓ 879ms | http |
| 88.80.150.82:8080 | ✓ 1498ms | 否 | 否 | ✓ 1871ms | ✓ 1645ms | https |
| 116.80.82.231:3172 | ✓ 1934ms | 否 | ✓ 1644ms | 否 | ✓ 1684ms | http |
| 162.248.165.72:1080 | ✓ 1166ms | ✓ 1797ms | ✓ 1873ms | 否 | 否 | http |
| 117.159.239.54:22222 | ✓ 879ms | ✓ 1067ms | ✓ 848ms | ✓ 1109ms | ✓ 873ms | http |
| 117.159.239.50:22222 | ✓ 844ms | ✓ 1075ms | ✓ 870ms | ✓ 1194ms | ✓ 877ms | http |
| 120.240.35.175:22222 | ✓ 897ms | ✓ 1187ms | ✓ 1015ms | ✓ 1173ms | 否 | http |
| 120.240.35.161:22222 | ✓ 981ms | ✓ 1248ms | ✓ 914ms | ✓ 1145ms | 否 | http |
| 106.14.205.114:483 | ✓ 1058ms | ✓ 1028ms | ✓ 1162ms | ✓ 1069ms | ✓ 840ms | http |
| 222.184.48.242:22222 | ✓ 953ms | ✓ 1169ms | ✓ 860ms | ✓ 1167ms | ✓ 954ms | http |
| 120.240.35.177:22222 | ✓ 945ms | ✓ 1246ms | ✓ 923ms | ✓ 1168ms | ✓ 958ms | http |
| 120.240.29.168:22222 | ✓ 948ms | ✓ 1239ms | ✓ 958ms | ✓ 1168ms | ✓ 977ms | http |
| 120.198.141.75:22222 | ✓ 928ms | ✓ 1207ms | ✓ 1013ms | ✓ 1142ms | ✓ 996ms | http |
| 113.59.32.163:22222 | ✓ 1041ms | ✓ 1332ms | ✓ 1030ms | ✓ 1295ms | ✓ 1068ms | http |
| 113.59.32.142:22222 | ✓ 1056ms | ✓ 1326ms | ✓ 1021ms | 否 | ✓ 949ms | http |
| 120.79.99.232:8099 | ✓ 1203ms | ✓ 1525ms | ✓ 1304ms | ✓ 1440ms | ✓ 1181ms | http |
| 106.14.203.63:3333 | 否 | ✓ 1831ms | ✓ 1756ms | ✓ 1373ms | 否 | http |
| 103.215.36.88:15412 | 否 | 否 | ✓ 1038ms | ✓ 1348ms | ✓ 1057ms | http |
| 27.73.57.47:10007 | ✓ 1614ms | 否 | ✓ 1757ms | ✓ 1863ms | 否 | http |
| 103.39.51.190:8080 | ✓ 1778ms | 否 | 否 | ✓ 1480ms | ✓ 1321ms | http |
| 183.249.5.109:22222 | ✓ 876ms | ✓ 964ms | ✓ 791ms | ✓ 957ms | ✓ 817ms | http |
| 120.238.159.230:22222 | ✓ 998ms | ✓ 1231ms | 否 | ✓ 1145ms | ✓ 943ms | http |
| 120.238.159.228:22222 | ✓ 989ms | ✓ 1200ms | ✓ 918ms | ✓ 1161ms | ✓ 879ms | http |
| 168.235.110.63:3128 | ✓ 1063ms | 否 | ✓ 1247ms | ✓ 1401ms | ✓ 1136ms | http |
| 222.184.48.251:22222 | ✓ 1463ms | ✓ 1773ms | ✓ 1297ms | ✓ 1623ms | ✓ 1188ms | http |
| 178.236.245.59:3128 | ✓ 1019ms | 否 | ✓ 1662ms | 否 | ✓ 1472ms | http |
| 120.238.159.229:22222 | ✓ 962ms | ✓ 1203ms | ✓ 981ms | ✓ 1123ms | ✓ 914ms | http |
| 61.72.221.94:3128 | ✓ 672ms | ✓ 1420ms | ✓ 1143ms | ✓ 1015ms | ✓ 1442ms | http |
| 121.128.121.54:3128 | ✓ 930ms | ✓ 1394ms | ✓ 1581ms | ✓ 1074ms | ✓ 823ms | http |
| 125.128.12.144:3128 | 否 | ✓ 1573ms | 否 | ✓ 1872ms | ✓ 993ms | http |
| 116.80.82.224:3172 | ✓ 1527ms | 否 | ✓ 1568ms | 否 | ✓ 1702ms | http |
| 115.231.181.40:8128 | ✓ 1925ms | 否 | ✓ 921ms | ✓ 1833ms | ✓ 963ms | http |
| 113.255.59.226:8080 | ✓ 1523ms | 否 | ✓ 1295ms | ✓ 1098ms | ✓ 1445ms | http |
| 45.186.6.104:3128 | ✓ 1632ms | ✓ 1942ms | ✓ 1756ms | 否 | 否 | http |
| 107.172.125.217:3128 | ✓ 757ms | 否 | ✓ 941ms | ✓ 1001ms | ✓ 750ms | http |
| 220.121.154.240:3128 | ✓ 1823ms | ✓ 1213ms | ✓ 1012ms | ✓ 988ms | ✓ 873ms | http |
| 103.215.36.88:17439 | ✓ 1151ms | ✓ 1350ms | 否 | ✓ 1131ms | 否 | http |
| 120.232.242.119:22222 | ✓ 910ms | 否 | ✓ 970ms | ✓ 1131ms | 否 | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1778ms | ✓ 1275ms | ✓ 1512ms | http |
| 178.236.245.17:3128 | ✓ 1374ms | ✓ 1897ms | ✓ 1601ms | 否 | ✓ 1694ms | http |
| 45.129.141.143:3128 | ✓ 1357ms | 否 | ✓ 1906ms | ✓ 1855ms | ✓ 1582ms | http |
| 61.72.221.194:3128 | ✓ 1689ms | 否 | 否 | ✓ 1642ms | ✓ 1394ms | http |
| 103.215.36.88:13809 | ✓ 1009ms | ✓ 1429ms | ✓ 1096ms | 否 | 否 | http |
| 163.44.126.97:3128 | 否 | 否 | ✓ 1939ms | ✓ 1704ms | ✓ 1128ms | http |
| 103.215.36.88:16350 | ✓ 864ms | 否 | ✓ 873ms | ✓ 1281ms | ✓ 915ms | http |
| 149.62.191.202:3128 | ✓ 1376ms | 否 | ✓ 1585ms | 否 | ✓ 1462ms | http |
| 45.136.198.40:3128 | ✓ 778ms | ✓ 1904ms | ✓ 1059ms | ✓ 1918ms | ✓ 1539ms | http |
| 192.71.213.85:9091 | ✓ 998ms | 否 | ✓ 764ms | ✓ 1865ms | 否 | http |
| 121.230.9.160:1080 | ✓ 1584ms | ✓ 1891ms | ✓ 1339ms | 否 | 否 | http |
| 39.104.201.40:7890 | 否 | ✓ 1215ms | 否 | ✓ 1507ms | ✓ 958ms | http |
| 120.240.35.176:22222 | ✓ 895ms | ✓ 1239ms | ✓ 972ms | ✓ 1143ms | ✓ 891ms | http |
| 103.69.84.106:3131 | ✓ 1528ms | 否 | ✓ 1352ms | ✓ 1479ms | ✓ 923ms | http |
| 121.230.8.245:1080 | ✓ 1351ms | ✓ 1347ms | ✓ 1360ms | 否 | 否 | http |
| 23.95.76.201:8443 | ✓ 672ms | 否 | ✓ 915ms | ✓ 1406ms | ✓ 1069ms | http |
| 180.125.216.109:8118 | ✓ 886ms | ✓ 1141ms | ✓ 887ms | 否 | ✓ 1129ms | http |
| 170.78.208.245:999 | ✓ 1493ms | 否 | ✓ 1585ms | 否 | ✓ 1541ms | http |
| 47.77.193.180:1080 | ✓ 1118ms | 否 | ✓ 356ms | ✓ 728ms | ✓ 599ms | http |
| 14.56.177.44:3128 | ✓ 1720ms | ✓ 1567ms | ✓ 1403ms | ✓ 1881ms | ✓ 814ms | http |
| 117.159.239.49:22222 | ✓ 849ms | ✓ 1174ms | ✓ 873ms | ✓ 1109ms | ✓ 856ms | http |
| 222.184.48.248:22222 | ✓ 1165ms | ✓ 1463ms | 否 | 否 | ✓ 1154ms | http |
| 103.215.36.88:10159 | ✓ 1007ms | ✓ 1368ms | ✓ 1185ms | ✓ 1374ms | 否 | http |
| 103.215.36.88:18867 | ✓ 1209ms | ✓ 1249ms | ✓ 1204ms | ✓ 1312ms | 否 | http |
| 103.215.36.88:18467 | ✓ 1237ms | ✓ 1263ms | ✓ 1116ms | ✓ 1350ms | ✓ 980ms | http |

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
