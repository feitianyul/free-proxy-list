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

最后更新：2026-04-18 11:37:12 UTC（2026-04-18 19:37:12 UTC+8）

**代理总数：75**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 75 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 38.59.240.157:12345 | ✓ 301ms | ✓ 826ms | 否 | ✓ 1071ms | ✓ 736ms | http |
| 1.231.81.166:3128 | ✓ 1492ms | ✓ 1449ms | ✓ 870ms | ✓ 887ms | ✓ 698ms | http |
| 218.108.131.186:17890 | ✓ 807ms | ✓ 1019ms | ✓ 917ms | ✓ 1325ms | ✓ 886ms | http |
| 223.84.151.86:30005 | ✓ 1164ms | ✓ 1251ms | ✓ 1018ms | ✓ 1375ms | ✓ 1179ms | http |
| 106.10.55.212:1121 | ✓ 1484ms | 否 | ✓ 1188ms | ✓ 1022ms | ✓ 1099ms | http |
| 157.230.178.216:8088 | 否 | 否 | ✓ 1273ms | ✓ 1708ms | ✓ 1344ms | http |
| 149.51.42.10:3128 | ✓ 532ms | ✓ 1584ms | 否 | ✓ 1455ms | 否 | http |
| 133.18.123.225:26021 | ✓ 538ms | 否 | ✓ 642ms | ✓ 840ms | ✓ 1209ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1145ms | ✓ 1248ms | ✓ 937ms | http |
| 45.76.207.177:40000 | ✓ 1823ms | 否 | 否 | ✓ 1881ms | ✓ 1430ms | http |
| 177.93.132.244:3128 | ✓ 785ms | 否 | ✓ 771ms | 否 | ✓ 1719ms | http |
| 113.160.132.26:8080 | ✓ 1445ms | 否 | ✓ 877ms | ✓ 1310ms | ✓ 951ms | http |
| 185.138.116.150:8080 | ✓ 819ms | 否 | 否 | ✓ 1679ms | ✓ 1398ms | http |
| 162.19.253.202:8443 | ✓ 826ms | 否 | ✓ 1519ms | 否 | ✓ 1925ms | http |
| 45.140.147.82:1081 | ✓ 1481ms | ✓ 1412ms | ✓ 586ms | ✓ 1742ms | 否 | http |
| 120.92.212.16:7890 | ✓ 930ms | 否 | 否 | ✓ 1171ms | ✓ 1407ms | http |
| 116.58.161.203:26021 | ✓ 1942ms | ✓ 1189ms | ✓ 1258ms | 否 | ✓ 1803ms | http |
| 120.92.108.86:7890 | ✓ 1607ms | 否 | ✓ 1714ms | ✓ 1876ms | 否 | http |
| 208.87.243.199:7878 | 否 | 否 | ✓ 1457ms | ✓ 1339ms | ✓ 880ms | http |
| 188.246.224.49:7890 | ✓ 738ms | 否 | ✓ 683ms | ✓ 1896ms | ✓ 1369ms | http |
| 130.61.30.221:8080 | ✓ 776ms | ✓ 1790ms | ✓ 1807ms | 否 | ✓ 1837ms | http |
| 159.89.191.221:3128 | ✓ 1952ms | ✓ 1448ms | ✓ 286ms | 否 | 否 | http |
| 91.99.15.45:2095 | ✓ 912ms | 否 | ✓ 1355ms | ✓ 1922ms | 否 | http |
| 212.58.132.5:8888 | ✓ 1787ms | 否 | ✓ 1464ms | 否 | ✓ 1201ms | http |
| 120.92.212.16:8890 | ✓ 1966ms | 否 | 否 | ✓ 1581ms | ✓ 1874ms | http |
| 82.148.18.242:443 | ✓ 964ms | 否 | ✓ 1750ms | 否 | ✓ 1928ms | http |
| 20.205.16.149:3128 | ✓ 661ms | ✓ 1205ms | ✓ 967ms | ✓ 828ms | 否 | http |
| 213.32.85.26:3128 | ✓ 690ms | ✓ 1776ms | ✓ 1359ms | 否 | 否 | http |
| 149.51.42.10:8080 | ✓ 1287ms | ✓ 1436ms | 否 | ✓ 1443ms | 否 | http |
| 103.138.70.165:3129 | ✓ 1283ms | 否 | ✓ 1406ms | ✓ 1653ms | ✓ 1397ms | http |
| 152.32.132.190:7890 | ✓ 1457ms | ✓ 1526ms | 否 | ✓ 844ms | ✓ 653ms | http |
| 168.144.75.9:3128 | ✓ 1664ms | 否 | ✓ 1927ms | ✓ 1990ms | ✓ 1577ms | http |
| 42.101.8.101:8888 | 否 | ✓ 1495ms | ✓ 1228ms | ✓ 1524ms | ✓ 1198ms | http |
| 84.47.150.125:1080 | ✓ 882ms | 否 | 否 | ✓ 1892ms | ✓ 1475ms | http |
| 62.113.119.14:8080 | ✓ 1291ms | ✓ 1926ms | ✓ 1004ms | 否 | 否 | http |
| 36.141.21.200:7890 | ✓ 915ms | ✓ 1213ms | ✓ 1000ms | ✓ 1297ms | 否 | http |
| 164.163.40.110:10000 | ✓ 1276ms | 否 | ✓ 1236ms | 否 | ✓ 1896ms | http |
| 59.46.216.131:30001 | ✓ 1085ms | 否 | ✓ 1232ms | ✓ 1466ms | 否 | http |
| 195.26.224.49:3128 | ✓ 1470ms | 否 | ✓ 567ms | ✓ 1950ms | 否 | http |
| 103.113.70.189:1082 | ✓ 1418ms | ✓ 1937ms | ✓ 1142ms | ✓ 1163ms | ✓ 1043ms | http |
| 154.193.243.86:40000 | ✓ 1592ms | ✓ 1448ms | ✓ 1799ms | ✓ 1478ms | ✓ 1077ms | http |
| 81.30.156.115:8080 | ✓ 1098ms | ✓ 1501ms | ✓ 1411ms | 否 | ✓ 1391ms | http |
| 43.132.188.134:443 | ✓ 834ms | 否 | ✓ 633ms | 否 | ✓ 1450ms | http |
| 116.80.60.44:7777 | 否 | 否 | ✓ 1477ms | ✓ 1819ms | ✓ 1671ms | http |
| 210.223.44.230:3128 | ✓ 632ms | 否 | ✓ 602ms | ✓ 1662ms | ✓ 977ms | http |
| 117.236.124.166:3128 | ✓ 1056ms | 否 | ✓ 1153ms | 否 | ✓ 1658ms | http |
| 113.176.92.71:3128 | ✓ 1937ms | ✓ 1307ms | ✓ 1555ms | ✓ 1530ms | ✓ 1263ms | http |
| 121.230.8.91:1080 | ✓ 1055ms | ✓ 1477ms | ✓ 971ms | 否 | ✓ 1173ms | http |
| 103.85.113.66:9999 | ✓ 655ms | ✓ 1513ms | ✓ 1702ms | 否 | ✓ 1755ms | http |
| 201.144.25.226:3128 | ✓ 1761ms | ✓ 1165ms | ✓ 988ms | ✓ 1299ms | ✓ 1091ms | http |
| 114.237.77.255:1080 | ✓ 977ms | ✓ 1775ms | ✓ 948ms | ✓ 1307ms | ✓ 1050ms | http |
| 121.230.8.62:1080 | ✓ 1312ms | ✓ 1673ms | ✓ 1199ms | 否 | 否 | http |
| 193.23.194.147:3128 | ✓ 1508ms | 否 | ✓ 1404ms | 否 | ✓ 1708ms | http |
| 125.76.214.178:8091 | ✓ 1302ms | ✓ 1230ms | ✓ 1036ms | ✓ 1478ms | ✓ 1217ms | http |
| 103.113.70.189:1081 | ✓ 935ms | 否 | ✓ 411ms | 否 | ✓ 1118ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1895ms | ✓ 1139ms | ✓ 936ms | http |
| 101.32.243.189:80 | ✓ 1430ms | 否 | ✓ 1482ms | ✓ 1861ms | ✓ 1270ms | http |
| 91.193.240.157:9877 | ✓ 1257ms | 否 | ✓ 1760ms | 否 | ✓ 1941ms | http |
| 178.140.10.58:1080 | ✓ 704ms | 否 | ✓ 1558ms | 否 | ✓ 1820ms | http |
| 117.122.240.82:3338 | ✓ 1258ms | ✓ 1340ms | ✓ 867ms | 否 | ✓ 920ms | http |
| 94.131.118.39:1082 | ✓ 1121ms | 否 | ✓ 1025ms | 否 | ✓ 1494ms | http |
| 94.131.118.39:1081 | ✓ 1140ms | 否 | ✓ 1004ms | 否 | ✓ 1462ms | http |
| 43.167.192.85:8080 | ✓ 994ms | 否 | ✓ 1490ms | 否 | ✓ 908ms | http |
| 14.247.76.52:8080 | 否 | 否 | ✓ 1734ms | ✓ 1787ms | ✓ 978ms | http |
| 38.180.2.107:3128 | ✓ 1650ms | ✓ 1825ms | ✓ 1947ms | 否 | ✓ 1785ms | http |
| 8.219.195.129:1080 | ✓ 1486ms | 否 | ✓ 963ms | ✓ 1049ms | ✓ 828ms | http |
| 143.198.211.194:8080 | ✓ 1376ms | 否 | ✓ 887ms | ✓ 1069ms | ✓ 848ms | http |
| 144.31.27.49:1080 | ✓ 1031ms | 否 | ✓ 1921ms | 否 | ✓ 1920ms | http |
| 45.140.147.155:1082 | ✓ 630ms | ✓ 1312ms | ✓ 686ms | ✓ 1676ms | 否 | http |
| 91.233.223.147:3128 | ✓ 1332ms | 否 | ✓ 921ms | 否 | ✓ 1636ms | http |
| 161.97.184.191:8080 | ✓ 1234ms | ✓ 1751ms | 否 | 否 | ✓ 1787ms | http |
| 45.140.147.155:1081 | ✓ 650ms | ✓ 1402ms | ✓ 1282ms | 否 | 否 | http |
| 20.127.128.70:8080 | ✓ 1888ms | 否 | ✓ 1345ms | 否 | ✓ 1870ms | http |
| 121.230.9.160:1080 | ✓ 1616ms | ✓ 1302ms | ✓ 1237ms | ✓ 1341ms | ✓ 1135ms | http |
| 61.52.131.172:8443 | 否 | ✓ 1086ms | ✓ 915ms | ✓ 1115ms | ✓ 997ms | http |

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
