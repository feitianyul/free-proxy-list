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

最后更新：2026-03-19 09:38:29 UTC（2026-03-19 17:38:29 UTC+8）

**代理总数：86**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 85 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 86 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1583ms | 否 | ✓ 807ms | ✓ 978ms | ✓ 1088ms | http |
| 202.155.12.161:443 | ✓ 1586ms | 否 | ✓ 1218ms | ✓ 1179ms | ✓ 1252ms | http |
| 147.161.239.240:8800 | ✓ 1365ms | 否 | ✓ 1153ms | ✓ 1788ms | ✓ 1327ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1604ms | ✓ 1479ms | ✓ 1406ms | ✓ 1007ms | http |
| 85.198.96.242:3128 | ✓ 1401ms | 否 | 否 | ✓ 1831ms | ✓ 1377ms | http |
| 115.231.181.40:8128 | ✓ 1766ms | 否 | 否 | ✓ 1443ms | ✓ 1631ms | http |
| 45.167.124.52:8080 | ✓ 842ms | 否 | 否 | ✓ 1661ms | ✓ 1374ms | http |
| 168.235.110.63:3128 | ✓ 330ms | ✓ 1093ms | ✓ 1794ms | ✓ 1163ms | ✓ 954ms | http |
| 137.220.150.152:6005 | 否 | 否 | ✓ 866ms | ✓ 1312ms | ✓ 1080ms | http |
| 45.125.67.37:443 | 否 | 否 | ✓ 1388ms | ✓ 1259ms | ✓ 1360ms | http |
| 137.220.151.110:6005 | 否 | 否 | ✓ 826ms | ✓ 1441ms | ✓ 1148ms | http |
| 222.184.48.251:22222 | ✓ 1701ms | ✓ 1697ms | ✓ 1269ms | ✓ 1635ms | 否 | http |
| 38.34.179.78:8448 | 否 | 否 | ✓ 1528ms | ✓ 1494ms | ✓ 986ms | http |
| 38.34.179.60:8450 | ✓ 1791ms | ✓ 1170ms | ✓ 340ms | ✓ 1021ms | ✓ 718ms | http |
| 219.117.204.211:7799 | ✓ 1771ms | 否 | ✓ 1602ms | ✓ 1049ms | ✓ 1947ms | http |
| 120.92.212.16:8890 | ✓ 953ms | ✓ 1170ms | ✓ 916ms | ✓ 1168ms | ✓ 1157ms | http |
| 120.92.212.16:7890 | ✓ 953ms | ✓ 1420ms | ✓ 1151ms | ✓ 1459ms | 否 | http |
| 183.249.5.117:22222 | ✓ 886ms | ✓ 979ms | ✓ 1000ms | ✓ 1068ms | 否 | http |
| 35.225.22.61:80 | ✓ 1139ms | ✓ 1309ms | ✓ 1221ms | 否 | 否 | http |
| 38.145.203.97:8448 | ✓ 915ms | 否 | ✓ 450ms | ✓ 919ms | ✓ 657ms | http |
| 174.138.24.77:1080 | ✓ 1490ms | 否 | ✓ 851ms | ✓ 1134ms | ✓ 957ms | http |
| 137.220.150.104:6005 | 否 | 否 | ✓ 1577ms | ✓ 1258ms | ✓ 946ms | http |
| 160.250.5.22:1 | ✓ 1692ms | 否 | ✓ 1419ms | ✓ 1531ms | ✓ 1061ms | http |
| 38.145.218.163:8451 | ✓ 395ms | 否 | ✓ 230ms | ✓ 832ms | ✓ 694ms | http |
| 138.124.53.25:7443 | ✓ 620ms | 否 | ✓ 1483ms | ✓ 1835ms | ✓ 1300ms | http |
| 212.192.12.90:6005 | ✓ 1249ms | 否 | ✓ 1773ms | ✓ 1206ms | ✓ 1391ms | http |
| 4.216.195.194:3128 | ✓ 640ms | 否 | ✓ 814ms | ✓ 1248ms | ✓ 805ms | http |
| 210.45.70.16:7895 | ✓ 1177ms | ✓ 1670ms | ✓ 1507ms | 否 | ✓ 1405ms | http |
| 209.126.10.139:3128 | ✓ 465ms | 否 | ✓ 862ms | ✓ 1393ms | 否 | http |
| 101.43.127.100:8877 | ✓ 765ms | ✓ 985ms | ✓ 841ms | ✓ 1064ms | ✓ 1067ms | http |
| 103.84.95.54:7890 | ✓ 754ms | 否 | ✓ 748ms | 否 | ✓ 1700ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 692ms | ✓ 1157ms | ✓ 835ms | http |
| 38.55.107.254:6005 | ✓ 1154ms | 否 | ✓ 1116ms | ✓ 1090ms | ✓ 1140ms | http |
| 38.55.107.137:6005 | ✓ 1258ms | 否 | ✓ 1163ms | ✓ 1089ms | ✓ 1158ms | http |
| 38.55.106.206:6005 | ✓ 1158ms | 否 | ✓ 1020ms | ✓ 1345ms | ✓ 1191ms | http |
| 38.55.106.208:6005 | ✓ 1211ms | 否 | ✓ 1016ms | ✓ 1041ms | 否 | http |
| 103.52.114.95:3128 | ✓ 1759ms | 否 | ✓ 1102ms | ✓ 1421ms | ✓ 1097ms | http |
| 45.93.29.147:6005 | ✓ 1710ms | ✓ 1292ms | 否 | 否 | ✓ 1320ms | http |
| 222.184.48.252:22222 | ✓ 987ms | ✓ 1995ms | 否 | ✓ 1801ms | ✓ 742ms | http |
| 103.113.70.189:1081 | ✓ 356ms | 否 | ✓ 464ms | ✓ 1230ms | ✓ 1783ms | http |
| 202.38.72.235:26001 | 否 | ✓ 1716ms | ✓ 1406ms | ✓ 1693ms | 否 | http |
| 59.46.216.131:30001 | ✓ 987ms | 否 | ✓ 1348ms | ✓ 1539ms | 否 | http |
| 137.220.150.170:6005 | ✓ 865ms | 否 | ✓ 1224ms | ✓ 1262ms | ✓ 975ms | http |
| 45.136.131.54:8448 | 否 | 否 | ✓ 1373ms | ✓ 1544ms | ✓ 1443ms | http |
| 34.101.184.164:3128 | ✓ 1792ms | 否 | ✓ 1541ms | ✓ 1722ms | ✓ 1223ms | http |
| 133.242.138.34:8100 | ✓ 1731ms | ✓ 1673ms | ✓ 1387ms | ✓ 1063ms | ✓ 839ms | http |
| 185.115.74.185:8080 | ✓ 1504ms | ✓ 1876ms | ✓ 1711ms | 否 | 否 | http |
| 38.145.220.198:8448 | 否 | 否 | ✓ 885ms | ✓ 1032ms | ✓ 738ms | http |
| 38.180.2.107:3128 | ✓ 1395ms | 否 | ✓ 1903ms | 否 | ✓ 1788ms | http |
| 88.80.150.82:8080 | ✓ 1299ms | 否 | ✓ 1962ms | 否 | ✓ 1722ms | https |
| 1.231.81.166:3128 | ✓ 1716ms | 否 | ✓ 1179ms | ✓ 1427ms | ✓ 1228ms | http |
| 103.139.138.194:3128 | ✓ 1838ms | 否 | ✓ 1235ms | ✓ 1490ms | ✓ 1228ms | http |
| 38.34.179.29:8452 | ✓ 709ms | ✓ 1057ms | ✓ 481ms | ✓ 993ms | ✓ 731ms | http |
| 38.34.179.61:8445 | ✓ 716ms | ✓ 1422ms | ✓ 1435ms | 否 | 否 | http |
| 103.82.93.219:3128 | ✓ 1684ms | 否 | ✓ 970ms | ✓ 1615ms | ✓ 1347ms | http |
| 45.136.198.40:3128 | ✓ 799ms | ✓ 1841ms | 否 | 否 | ✓ 1799ms | http |
| 149.62.191.202:3128 | 否 | 否 | ✓ 1578ms | ✓ 1974ms | ✓ 1475ms | http |
| 8.219.97.248:80 | ✓ 1262ms | 否 | ✓ 1163ms | 否 | ✓ 1848ms | http |
| 120.55.163.237:10086 | ✓ 765ms | ✓ 891ms | ✓ 832ms | ✓ 1013ms | ✓ 853ms | http |
| 121.237.181.137:8888 | ✓ 924ms | ✓ 1086ms | ✓ 887ms | ✓ 1355ms | 否 | http |
| 106.117.208.101:7890 | ✓ 1372ms | ✓ 1858ms | ✓ 1784ms | ✓ 1202ms | 否 | http |
| 101.47.73.135:3128 | ✓ 1727ms | 否 | ✓ 1358ms | ✓ 1434ms | ✓ 1669ms | http |
| 84.247.149.172:3128 | ✓ 1245ms | 否 | 否 | ✓ 1523ms | ✓ 1217ms | http |
| 47.77.193.180:1080 | ✓ 326ms | ✓ 1364ms | ✓ 312ms | ✓ 802ms | ✓ 611ms | http |
| 207.254.71.62:8088 | ✓ 979ms | 否 | ✓ 851ms | ✓ 1755ms | ✓ 1382ms | http |
| 223.16.170.103:80 | 否 | 否 | ✓ 934ms | ✓ 1162ms | ✓ 1201ms | http |
| 24.144.86.173:1080 | ✓ 1364ms | 否 | ✓ 1239ms | ✓ 890ms | ✓ 649ms | http |
| 45.149.92.147:5001 | ✓ 738ms | 否 | ✓ 911ms | 否 | ✓ 726ms | http |
| 5.102.109.41:999 | 否 | 否 | ✓ 1106ms | ✓ 1378ms | ✓ 1031ms | http |
| 113.176.92.71:3128 | 否 | 否 | ✓ 1555ms | ✓ 1454ms | ✓ 1055ms | http |
| 43.165.195.107:3128 | ✓ 1663ms | 否 | ✓ 1183ms | ✓ 1303ms | ✓ 1039ms | http |
| 103.39.51.190:8080 | ✓ 1674ms | 否 | 否 | ✓ 1727ms | ✓ 1451ms | http |
| 85.208.108.43:2094 | ✓ 1299ms | 否 | ✓ 597ms | ✓ 1126ms | ✓ 742ms | http |
| 121.230.8.214:1080 | 否 | ✓ 1741ms | ✓ 1464ms | 否 | ✓ 1472ms | http |
| 62.113.119.14:8080 | ✓ 729ms | 否 | ✓ 1485ms | ✓ 1728ms | ✓ 1340ms | http |
| 114.237.77.231:1080 | ✓ 948ms | 否 | ✓ 981ms | 否 | ✓ 1493ms | http |
| 106.14.203.63:3333 | ✓ 1629ms | ✓ 1941ms | ✓ 1652ms | ✓ 1892ms | 否 | http |
| 201.144.20.238:3128 | ✓ 696ms | 否 | ✓ 961ms | ✓ 1159ms | 否 | http |
| 202.141.161.53:30001 | 否 | ✓ 1426ms | ✓ 1221ms | ✓ 1314ms | ✓ 1092ms | http |
| 172.212.68.37:3128 | ✓ 1287ms | 否 | ✓ 1734ms | ✓ 1261ms | ✓ 1360ms | http |
| 20.120.225.109:3128 | ✓ 904ms | ✓ 1863ms | ✓ 1764ms | ✓ 1726ms | 否 | http |
| 38.145.203.19:8447 | ✓ 1765ms | 否 | ✓ 1852ms | ✓ 1207ms | ✓ 1997ms | http |
| 121.230.9.19:1080 | ✓ 1136ms | 否 | ✓ 980ms | ✓ 1618ms | 否 | http |
| 45.168.238.193:8443 | ✓ 977ms | ✓ 1118ms | ✓ 236ms | ✓ 1192ms | ✓ 934ms | http |
| 146.56.182.165:3128 | ✓ 1582ms | 否 | 否 | ✓ 1752ms | ✓ 1085ms | http |
| 45.136.130.197:8452 | ✓ 954ms | ✓ 1069ms | ✓ 1862ms | ✓ 845ms | ✓ 841ms | http |

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
