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

最后更新：2026-05-16 10:41:16 UTC（2026-05-16 18:41:16 UTC+8）

**代理总数：76**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 107.175.85.198:1080 | ✓ 298ms | 否 | ✓ 864ms | ✓ 1527ms | ✓ 1011ms | http |
| 103.134.85.167:3128 | ✓ 1692ms | 否 | ✓ 1127ms | ✓ 1468ms | ✓ 1219ms | http |
| 185.200.188.234:10001 | ✓ 1262ms | 否 | ✓ 1774ms | 否 | ✓ 1698ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1710ms | ✓ 1853ms | ✓ 1910ms | 否 | http |
| 129.80.238.83:444 | ✓ 945ms | ✓ 1691ms | ✓ 1274ms | ✓ 879ms | ✓ 874ms | http |
| 129.80.217.21:444 | ✓ 868ms | 否 | ✓ 840ms | ✓ 862ms | ✓ 1030ms | http |
| 212.58.132.5:8888 | ✓ 998ms | 否 | ✓ 995ms | ✓ 1420ms | ✓ 1148ms | http |
| 45.125.67.37:8443 | ✓ 1090ms | 否 | ✓ 1729ms | ✓ 1542ms | ✓ 1314ms | http |
| 43.156.90.221:10808 | ✓ 1683ms | 否 | ✓ 1719ms | ✓ 1263ms | ✓ 1008ms | http |
| 82.114.228.67:1080 | 否 | 否 | ✓ 1107ms | ✓ 1929ms | ✓ 1452ms | http |
| 1.231.81.166:3128 | 否 | 否 | ✓ 1701ms | ✓ 1306ms | ✓ 1308ms | http |
| 5.252.33.13:2025 | ✓ 1295ms | 否 | ✓ 1277ms | 否 | ✓ 1601ms | http |
| 185.21.15.206:3128 | 否 | 否 | ✓ 638ms | ✓ 1684ms | ✓ 1484ms | http |
| 137.59.47.73:3128 | ✓ 1683ms | 否 | ✓ 1490ms | ✓ 1732ms | 否 | http |
| 168.222.254.136:8888 | 否 | ✓ 1764ms | ✓ 1590ms | 否 | ✓ 1296ms | http |
| 158.160.215.167:8124 | ✓ 699ms | ✓ 1730ms | ✓ 1774ms | 否 | 否 | http |
| 91.242.229.129:8092 | ✓ 594ms | ✓ 1674ms | ✓ 637ms | ✓ 1656ms | 否 | http |
| 190.93.224.32:999 | ✓ 1110ms | 否 | ✓ 947ms | ✓ 1946ms | ✓ 1701ms | http |
| 5.129.248.58:3128 | ✓ 1001ms | ✓ 1609ms | ✓ 1863ms | 否 | 否 | http |
| 34.71.229.255:3128 | ✓ 261ms | ✓ 1345ms | ✓ 1015ms | ✓ 1394ms | ✓ 859ms | http |
| 103.21.220.141:3128 | ✓ 837ms | 否 | ✓ 823ms | ✓ 1054ms | ✓ 841ms | http |
| 103.157.200.126:3128 | ✓ 1747ms | 否 | ✓ 1914ms | 否 | ✓ 1896ms | http |
| 147.45.186.28:3128 | ✓ 840ms | ✓ 1781ms | ✓ 843ms | ✓ 1686ms | ✓ 1628ms | http |
| 148.230.4.241:999 | ✓ 807ms | ✓ 1813ms | ✓ 1323ms | ✓ 1550ms | ✓ 1298ms | http |
| 128.199.121.61:9090 | ✓ 939ms | 否 | ✓ 1512ms | ✓ 1373ms | ✓ 1066ms | http |
| 43.167.192.85:8080 | ✓ 1462ms | ✓ 1290ms | ✓ 750ms | ✓ 1146ms | ✓ 865ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1598ms | 否 | ✓ 1421ms | ✓ 1232ms | http |
| 210.223.44.230:3128 | 否 | 否 | ✓ 1421ms | ✓ 1811ms | ✓ 1803ms | http |
| 45.59.122.132:80 | 否 | ✓ 1662ms | 否 | ✓ 1290ms | ✓ 1242ms | http |
| 152.70.91.193:40000 | ✓ 1932ms | 否 | 否 | ✓ 1690ms | ✓ 1909ms | http |
| 103.147.152.12:1080 | ✓ 847ms | 否 | ✓ 1699ms | 否 | ✓ 1090ms | http |
| 34.96.238.40:8080 | ✓ 1130ms | 否 | ✓ 1220ms | ✓ 1232ms | ✓ 1315ms | http |
| 84.47.150.125:1080 | ✓ 1542ms | 否 | ✓ 1678ms | 否 | ✓ 1661ms | http |
| 34.101.184.164:3128 | ✓ 1703ms | 否 | ✓ 1409ms | ✓ 1565ms | ✓ 1203ms | http |
| 168.110.52.228:3128 | ✓ 784ms | ✓ 1662ms | 否 | 否 | ✓ 831ms | http |
| 218.108.131.186:17890 | ✓ 1007ms | ✓ 1303ms | ✓ 1078ms | ✓ 1351ms | ✓ 1061ms | http |
| 217.174.244.117:3129 | ✓ 1088ms | ✓ 1867ms | ✓ 1320ms | 否 | ✓ 1552ms | http |
| 160.238.65.6:3128 | ✓ 1252ms | ✓ 1616ms | 否 | 否 | ✓ 1414ms | http |
| 185.230.191.240:3128 | 否 | ✓ 1640ms | ✓ 1783ms | ✓ 1824ms | 否 | http |
| 91.217.81.131:1080 | ✓ 670ms | 否 | ✓ 1030ms | 否 | ✓ 1375ms | http |
| 166.88.55.83:7890 | ✓ 811ms | ✓ 1539ms | ✓ 815ms | ✓ 1079ms | ✓ 923ms | http |
| 103.147.152.12:1095 | ✓ 1726ms | 否 | ✓ 689ms | 否 | ✓ 1129ms | http |
| 47.74.226.8:5001 | 否 | ✓ 1655ms | ✓ 1325ms | ✓ 1540ms | 否 | http |
| 8.219.97.248:80 | ✓ 1751ms | 否 | ✓ 1992ms | ✓ 1727ms | 否 | http |
| 128.199.254.13:9090 | 否 | 否 | ✓ 1469ms | ✓ 1314ms | ✓ 1063ms | http |
| 128.199.114.189:9090 | 否 | 否 | ✓ 1368ms | ✓ 1333ms | ✓ 1047ms | http |
| 152.42.170.187:9090 | 否 | 否 | ✓ 1243ms | ✓ 1317ms | ✓ 1059ms | http |
| 128.199.116.219:9090 | ✓ 1575ms | 否 | ✓ 972ms | ✓ 1344ms | ✓ 1084ms | http |
| 160.238.65.5:3128 | ✓ 1928ms | ✓ 1697ms | 否 | 否 | ✓ 1545ms | http |
| 160.238.65.2:3128 | ✓ 1928ms | 否 | 否 | ✓ 1280ms | ✓ 1606ms | http |
| 152.42.177.32:8888 | ✓ 1181ms | 否 | ✓ 1197ms | ✓ 1562ms | 否 | http |
| 160.238.65.3:3128 | ✓ 693ms | ✓ 1537ms | 否 | ✓ 1824ms | 否 | http |
| 128.199.113.85:9090 | ✓ 925ms | 否 | ✓ 981ms | ✓ 1309ms | ✓ 1050ms | http |
| 146.190.80.158:9090 | ✓ 1297ms | 否 | ✓ 946ms | ✓ 1343ms | ✓ 1054ms | http |
| 121.130.177.28:8888 | 否 | ✓ 1676ms | ✓ 1749ms | ✓ 1676ms | ✓ 1487ms | http |
| 8.154.21.175:3128 | ✓ 1627ms | ✓ 1929ms | ✓ 1869ms | ✓ 1316ms | ✓ 1699ms | http |
| 3.101.133.120:80 | ✓ 410ms | 否 | ✓ 1538ms | ✓ 1415ms | ✓ 939ms | http |
| 57.129.144.178:40000 | ✓ 1244ms | ✓ 1619ms | ✓ 1032ms | 否 | ✓ 1015ms | http |
| 200.174.198.32:8888 | ✓ 1170ms | 否 | ✓ 920ms | 否 | ✓ 1713ms | http |
| 157.0.142.246:10057 | ✓ 1575ms | ✓ 1514ms | ✓ 1203ms | ✓ 1675ms | ✓ 1239ms | http |
| 59.46.216.131:30001 | ✓ 1274ms | ✓ 1630ms | ✓ 1228ms | ✓ 1579ms | 否 | http |
| 3.15.187.17:1080 | ✓ 277ms | ✓ 1884ms | ✓ 1526ms | ✓ 1338ms | ✓ 1125ms | http |
| 45.129.141.143:3128 | ✓ 631ms | ✓ 1686ms | ✓ 1789ms | ✓ 1790ms | ✓ 1617ms | http |
| 104.248.151.93:9090 | 否 | 否 | ✓ 985ms | ✓ 1301ms | ✓ 1022ms | http |
| 106.10.55.212:1121 | ✓ 1634ms | 否 | ✓ 1543ms | ✓ 1587ms | ✓ 1151ms | http |
| 103.195.142.250:8180 | ✓ 1897ms | 否 | ✓ 1850ms | ✓ 1958ms | 否 | http |
| 86.104.74.110:1081 | ✓ 1027ms | ✓ 1314ms | ✓ 880ms | 否 | ✓ 1558ms | http |
| 86.104.74.110:1082 | ✓ 1078ms | ✓ 1219ms | ✓ 977ms | 否 | ✓ 1540ms | http |
| 45.186.6.104:3128 | ✓ 1467ms | ✓ 1595ms | ✓ 1894ms | 否 | 否 | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1002ms | ✓ 1330ms | ✓ 1071ms | http |
| 20.164.75.153:8080 | ✓ 955ms | 否 | ✓ 1153ms | 否 | ✓ 1840ms | http |
| 64.181.240.152:3128 | ✓ 1784ms | 否 | ✓ 1950ms | 否 | ✓ 1847ms | http |
| 77.110.119.136:3128 | ✓ 194ms | 否 | ✓ 186ms | ✓ 916ms | ✓ 982ms | http |
| 180.103.19.117:1080 | ✓ 1402ms | ✓ 1602ms | 否 | 否 | ✓ 1319ms | http |
| 61.52.131.172:8443 | ✓ 1087ms | ✓ 1338ms | ✓ 1174ms | ✓ 1526ms | ✓ 1248ms | http |
| 114.214.170.41:27890 | ✓ 1258ms | ✓ 1629ms | ✓ 1496ms | ✓ 1571ms | ✓ 1276ms | http |

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
