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

最后更新：2026-03-02 21:37:57 UTC（2026-03-03 05:37:57 UTC+8）

**代理总数：77**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 162.240.154.26:3128 | ✓ 830ms | 否 | ✓ 1338ms | ✓ 1326ms | ✓ 1368ms | http |
| 45.88.0.98:3128 | ✓ 897ms | ✓ 1798ms | 否 | 否 | ✓ 1738ms | http |
| 45.88.0.99:3128 | ✓ 1530ms | ✓ 1543ms | 否 | ✓ 1899ms | ✓ 1468ms | http |
| 45.140.147.82:1081 | ✓ 446ms | ✓ 1475ms | ✓ 839ms | ✓ 1896ms | ✓ 1273ms | http |
| 211.171.114.154:3128 | ✓ 1753ms | ✓ 1187ms | 否 | ✓ 1432ms | ✓ 1075ms | http |
| 35.234.17.221:8080 | ✓ 950ms | ✓ 1398ms | ✓ 1316ms | 否 | ✓ 1337ms | http |
| 217.76.245.80:999 | ✓ 958ms | ✓ 1546ms | ✓ 1807ms | ✓ 1701ms | ✓ 1346ms | http |
| 103.84.95.54:7890 | ✓ 916ms | 否 | 否 | ✓ 943ms | ✓ 1046ms | http |
| 95.85.252.153:21064 | ✓ 987ms | ✓ 1526ms | ✓ 1595ms | 否 | 否 | http |
| 115.76.5.32:10010 | 否 | 否 | ✓ 1494ms | ✓ 1911ms | ✓ 1653ms | http |
| 138.124.53.25:7443 | ✓ 797ms | 否 | ✓ 1398ms | ✓ 1720ms | ✓ 1246ms | http |
| 205.209.118.30:3138 | ✓ 598ms | ✓ 976ms | 否 | 否 | ✓ 883ms | http |
| 35.225.22.61:80 | ✓ 1069ms | 否 | ✓ 800ms | ✓ 1047ms | ✓ 737ms | http |
| 91.99.99.83:9000 | 否 | ✓ 1781ms | 否 | ✓ 1800ms | ✓ 1732ms | http |
| 81.70.169.194:80 | ✓ 1406ms | ✓ 1357ms | ✓ 1074ms | ✓ 1517ms | ✓ 1064ms | http |
| 101.43.255.96:80 | ✓ 1200ms | ✓ 1602ms | ✓ 1255ms | ✓ 1407ms | ✓ 1811ms | http |
| 150.107.140.238:3128 | ✓ 1648ms | 否 | ✓ 955ms | ✓ 1306ms | ✓ 1981ms | http |
| 74.48.78.224:2080 | ✓ 358ms | ✓ 873ms | ✓ 424ms | ✓ 945ms | ✓ 733ms | http |
| 121.128.121.54:3128 | ✓ 1706ms | 否 | ✓ 728ms | ✓ 1053ms | ✓ 856ms | http |
| 14.56.107.244:3128 | 否 | ✓ 1552ms | ✓ 742ms | ✓ 1945ms | ✓ 1280ms | http |
| 115.76.5.32:10005 | ✓ 1947ms | 否 | ✓ 1528ms | 否 | ✓ 1518ms | http |
| 14.56.177.44:3128 | ✓ 769ms | ✓ 1131ms | 否 | ✓ 1300ms | ✓ 894ms | http |
| 146.190.232.76:3128 | ✓ 951ms | 否 | ✓ 1969ms | 否 | ✓ 1352ms | http |
| 115.231.181.40:8128 | ✓ 1922ms | 否 | ✓ 997ms | ✓ 1586ms | 否 | http |
| 160.238.65.6:3128 | 否 | 否 | ✓ 513ms | ✓ 1351ms | ✓ 1052ms | http |
| 91.238.104.172:2024 | ✓ 1938ms | 否 | ✓ 1827ms | 否 | ✓ 1317ms | http |
| 45.140.147.82:1082 | ✓ 769ms | ✓ 1372ms | ✓ 866ms | ✓ 1274ms | ✓ 1259ms | http |
| 45.88.0.111:3128 | ✓ 1908ms | ✓ 1600ms | ✓ 1382ms | 否 | 否 | http |
| 91.238.104.171:2023 | ✓ 1195ms | 否 | ✓ 1675ms | 否 | ✓ 1644ms | http |
| 5.75.196.26:40000 | ✓ 508ms | ✓ 1468ms | ✓ 1996ms | 否 | 否 | http |
| 222.228.171.92:8080 | ✓ 1587ms | ✓ 1853ms | 否 | 否 | ✓ 1155ms | http |
| 115.76.5.32:10007 | ✓ 1657ms | 否 | 否 | ✓ 1953ms | ✓ 1546ms | http |
| 115.76.5.32:10009 | 否 | 否 | ✓ 1848ms | ✓ 1841ms | ✓ 1647ms | http |
| 45.88.0.113:3128 | ✓ 518ms | ✓ 1422ms | 否 | 否 | ✓ 1056ms | http |
| 83.219.250.8:62920 | ✓ 1008ms | 否 | ✓ 1539ms | 否 | ✓ 1895ms | http |
| 45.88.0.117:3128 | 否 | 否 | ✓ 1905ms | ✓ 1366ms | ✓ 1571ms | http |
| 120.92.212.16:7890 | ✓ 1813ms | 否 | 否 | ✓ 1320ms | ✓ 1844ms | http |
| 180.103.19.117:1080 | ✓ 1086ms | 否 | 否 | ✓ 1544ms | ✓ 1102ms | http |
| 121.230.9.161:1080 | ✓ 1576ms | ✓ 1548ms | ✓ 1248ms | ✓ 1522ms | ✓ 1474ms | http |
| 142.171.131.38:7890 | ✓ 1452ms | ✓ 844ms | ✓ 954ms | 否 | ✓ 1930ms | http |
| 34.101.184.164:3128 | ✓ 1719ms | 否 | ✓ 1528ms | 否 | ✓ 1581ms | http |
| 101.32.244.83:8080 | ✓ 1447ms | ✓ 1826ms | ✓ 1022ms | ✓ 1470ms | ✓ 1326ms | http |
| 121.43.196.210:8222 | ✓ 1024ms | ✓ 1136ms | ✓ 907ms | ✓ 1202ms | ✓ 962ms | http |
| 121.43.196.213:8222 | ✓ 1010ms | ✓ 1114ms | ✓ 949ms | ✓ 1214ms | ✓ 944ms | http |
| 114.55.226.123:10086 | ✓ 1076ms | ✓ 1532ms | ✓ 1166ms | ✓ 1363ms | ✓ 1098ms | http |
| 45.88.0.115:3128 | ✓ 628ms | 否 | ✓ 850ms | ✓ 1338ms | ✓ 1453ms | http |
| 45.88.0.116:3128 | ✓ 614ms | 否 | ✓ 1501ms | ✓ 1705ms | ✓ 1778ms | http |
| 123.20.24.166:8118 | ✓ 1787ms | 否 | ✓ 1617ms | ✓ 1728ms | ✓ 1559ms | http |
| 160.238.65.4:3128 | ✓ 970ms | 否 | 否 | ✓ 1344ms | ✓ 1060ms | http |
| 158.160.215.167:8127 | ✓ 1007ms | 否 | ✓ 1920ms | ✓ 1683ms | ✓ 1521ms | http |
| 172.212.68.37:3128 | ✓ 396ms | ✓ 1617ms | ✓ 1463ms | ✓ 1132ms | ✓ 1293ms | http |
| 2.56.178.131:443 | ✓ 1465ms | 否 | ✓ 1892ms | 否 | ✓ 1892ms | http |
| 120.92.212.16:8890 | 否 | ✓ 1629ms | 否 | ✓ 1299ms | ✓ 1040ms | http |
| 107.174.133.10:3128 | ✓ 701ms | ✓ 873ms | ✓ 947ms | 否 | ✓ 1049ms | http |
| 45.136.198.40:3128 | ✓ 1274ms | ✓ 1879ms | ✓ 1685ms | 否 | ✓ 1714ms | http |
| 124.156.179.148:3128 | ✓ 760ms | ✓ 984ms | ✓ 766ms | ✓ 924ms | ✓ 738ms | http |
| 192.73.243.65:3128 | ✓ 812ms | ✓ 1040ms | ✓ 1887ms | ✓ 1058ms | ✓ 839ms | http |
| 157.245.194.13:8888 | ✓ 791ms | 否 | ✓ 901ms | ✓ 1122ms | ✓ 921ms | http |
| 104.248.25.131:3128 | ✓ 530ms | 否 | ✓ 1521ms | ✓ 1775ms | ✓ 1535ms | http |
| 103.215.36.88:15556 | ✓ 1106ms | ✓ 1319ms | ✓ 1142ms | 否 | ✓ 1022ms | http |
| 23.236.70.105:8888 | ✓ 524ms | ✓ 1031ms | ✓ 1097ms | ✓ 1103ms | 否 | http |
| 38.180.2.107:3128 | ✓ 1171ms | 否 | ✓ 1888ms | 否 | ✓ 1709ms | http |
| 103.39.51.190:8080 | ✓ 1837ms | 否 | ✓ 1733ms | ✓ 1702ms | ✓ 1618ms | http |
| 45.125.67.37:8443 | ✓ 1076ms | 否 | ✓ 1105ms | 否 | ✓ 1083ms | http |
| 188.166.208.168:9876 | ✓ 1409ms | 否 | ✓ 1133ms | ✓ 1132ms | ✓ 908ms | http |
| 209.126.10.139:3128 | ✓ 696ms | 否 | ✓ 1949ms | ✓ 1396ms | ✓ 944ms | http |
| 125.128.12.144:3128 | ✓ 1638ms | ✓ 1882ms | ✓ 1582ms | 否 | 否 | http |
| 61.72.110.54:3128 | ✓ 922ms | 否 | 否 | ✓ 1263ms | ✓ 871ms | http |
| 84.247.149.172:3128 | ✓ 865ms | 否 | ✓ 1683ms | ✓ 1398ms | ✓ 954ms | http |
| 199.68.217.2:3128 | ✓ 683ms | ✓ 1663ms | ✓ 786ms | ✓ 1090ms | ✓ 828ms | http |
| 115.76.5.32:10008 | ✓ 1826ms | 否 | 否 | ✓ 1907ms | ✓ 1474ms | http |
| 45.88.0.114:3128 | ✓ 514ms | ✓ 1389ms | 否 | 否 | ✓ 1077ms | http |
| 94.159.112.138:3129 | ✓ 1098ms | 否 | ✓ 1172ms | 否 | ✓ 1776ms | http |
| 103.131.19.42:8181 | ✓ 1529ms | 否 | ✓ 1729ms | ✓ 1537ms | ✓ 1484ms | http |
| 47.77.180.205:1080 | ✓ 1080ms | ✓ 1170ms | ✓ 662ms | ✓ 960ms | 否 | http |
| 113.255.59.226:8080 | ✓ 1260ms | 否 | ✓ 1156ms | 否 | ✓ 1261ms | http |
| 45.186.6.104:3128 | ✓ 1434ms | ✓ 1573ms | ✓ 1901ms | 否 | 否 | http |

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
