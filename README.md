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

最后更新：2026-04-23 23:43:03 UTC（2026-04-24 07:43:03 UTC+8）

**代理总数：64**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.85.51.197:1080 | 否 | ✓ 1120ms | ✓ 256ms | ✓ 914ms | ✓ 1108ms | http |
| 1.231.81.166:3128 | ✓ 865ms | ✓ 1460ms | ✓ 1127ms | ✓ 1205ms | ✓ 881ms | http |
| 46.101.95.183:8888 | ✓ 1360ms | ✓ 1762ms | ✓ 1563ms | ✓ 1785ms | ✓ 1179ms | http |
| 115.231.181.40:8128 | ✓ 1090ms | ✓ 1369ms | 否 | 否 | ✓ 1173ms | http |
| 212.58.132.5:8888 | ✓ 995ms | 否 | ✓ 1030ms | ✓ 1376ms | ✓ 1115ms | http |
| 113.160.132.26:8080 | 否 | 否 | ✓ 1511ms | ✓ 1476ms | ✓ 1146ms | http |
| 120.92.108.86:7890 | ✓ 1347ms | 否 | ✓ 1394ms | ✓ 1864ms | ✓ 1482ms | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1287ms | ✓ 1507ms | ✓ 1257ms | http |
| 84.47.150.125:1080 | 否 | 否 | ✓ 1715ms | ✓ 1957ms | ✓ 1620ms | http |
| 45.153.231.229:8080 | ✓ 823ms | 否 | ✓ 1590ms | ✓ 1971ms | 否 | http |
| 130.61.174.200:1080 | ✓ 1493ms | 否 | ✓ 810ms | ✓ 1396ms | ✓ 1285ms | http |
| 42.200.76.16:3888 | ✓ 1751ms | 否 | ✓ 874ms | ✓ 1080ms | ✓ 856ms | http |
| 218.108.131.186:17890 | ✓ 807ms | ✓ 995ms | ✓ 852ms | 否 | 否 | http |
| 120.92.212.16:8890 | ✓ 1432ms | ✓ 1378ms | ✓ 1129ms | 否 | ✓ 1203ms | http |
| 35.225.22.61:80 | ✓ 873ms | ✓ 1410ms | 否 | 否 | ✓ 999ms | http |
| 120.92.212.16:7890 | ✓ 1098ms | ✓ 1497ms | ✓ 1087ms | ✓ 1432ms | ✓ 1487ms | http |
| 20.164.75.153:8080 | ✓ 1217ms | 否 | ✓ 1154ms | 否 | ✓ 1997ms | http |
| 103.187.146.151:3128 | ✓ 1672ms | 否 | ✓ 964ms | ✓ 1282ms | ✓ 1082ms | http |
| 210.45.76.58:42992 | ✓ 1171ms | ✓ 1449ms | ✓ 1363ms | ✓ 1507ms | ✓ 1166ms | http |
| 82.148.18.242:443 | ✓ 1283ms | ✓ 1800ms | ✓ 1867ms | 否 | 否 | http |
| 38.180.2.107:3128 | ✓ 708ms | ✓ 1628ms | ✓ 1852ms | 否 | ✓ 1829ms | http |
| 42.101.8.101:8888 | ✓ 1272ms | 否 | ✓ 1452ms | ✓ 1572ms | 否 | http |
| 138.124.99.216:8888 | ✓ 785ms | 否 | ✓ 1385ms | ✓ 1950ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1159ms | ✓ 1498ms | ✓ 948ms | ✓ 1473ms | ✓ 1221ms | http |
| 91.99.15.45:2095 | ✓ 1161ms | 否 | ✓ 1435ms | ✓ 1529ms | ✓ 1798ms | http |
| 8.211.166.184:8081 | ✓ 654ms | 否 | ✓ 848ms | ✓ 1075ms | ✓ 815ms | http |
| 163.44.126.97:3128 | ✓ 790ms | ✓ 1909ms | 否 | ✓ 1216ms | ✓ 896ms | http |
| 45.140.147.82:1082 | ✓ 1076ms | 否 | ✓ 956ms | 否 | ✓ 1784ms | http |
| 121.138.61.78:8481 | ✓ 938ms | ✓ 1690ms | ✓ 1320ms | ✓ 1194ms | ✓ 982ms | http |
| 146.56.185.39:10900 | 否 | 否 | ✓ 1906ms | ✓ 1189ms | ✓ 986ms | http |
| 194.87.57.226:3128 | ✓ 665ms | ✓ 1626ms | ✓ 1438ms | 否 | ✓ 1972ms | http |
| 34.101.184.164:3128 | ✓ 1663ms | 否 | ✓ 1275ms | ✓ 1454ms | ✓ 1285ms | http |
| 103.184.99.194:8080 | ✓ 1687ms | 否 | ✓ 1898ms | ✓ 1776ms | ✓ 1720ms | http |
| 121.230.8.22:1080 | ✓ 1564ms | ✓ 1691ms | ✓ 1207ms | ✓ 1697ms | ✓ 1287ms | http |
| 168.110.52.228:3128 | ✓ 779ms | ✓ 1421ms | ✓ 1509ms | ✓ 1117ms | ✓ 946ms | http |
| 45.140.147.82:1081 | ✓ 607ms | 否 | ✓ 1281ms | ✓ 1533ms | 否 | http |
| 223.84.151.86:30005 | ✓ 1883ms | ✓ 1939ms | ✓ 1219ms | ✓ 1699ms | 否 | http |
| 188.166.230.157:3128 | ✓ 1118ms | 否 | ✓ 928ms | ✓ 1281ms | ✓ 1011ms | http |
| 208.87.243.199:7878 | 否 | 否 | ✓ 1100ms | ✓ 1339ms | ✓ 1338ms | http |
| 137.59.47.73:3128 | ✓ 1987ms | ✓ 1955ms | ✓ 1148ms | 否 | 否 | http |
| 2.83.243.148:7777 | ✓ 763ms | ✓ 1838ms | ✓ 1704ms | ✓ 1469ms | ✓ 1234ms | http |
| 91.193.240.157:9877 | ✓ 1015ms | 否 | ✓ 1622ms | 否 | ✓ 1602ms | http |
| 152.32.132.190:7890 | ✓ 1905ms | 否 | 否 | ✓ 1118ms | ✓ 1745ms | http |
| 45.76.207.177:40000 | 否 | 否 | ✓ 871ms | ✓ 1441ms | ✓ 1272ms | http |
| 201.144.20.238:3128 | ✓ 1275ms | ✓ 1453ms | ✓ 1205ms | ✓ 1594ms | ✓ 1125ms | http |
| 121.230.8.136:1080 | ✓ 1209ms | ✓ 1404ms | ✓ 1147ms | ✓ 1574ms | ✓ 1150ms | http |
| 45.140.147.155:1082 | ✓ 1691ms | ✓ 1396ms | ✓ 851ms | 否 | 否 | http |
| 194.147.90.23:3128 | ✓ 692ms | ✓ 1676ms | ✓ 767ms | 否 | ✓ 1542ms | http |
| 202.129.206.239:3128 | ✓ 1362ms | 否 | ✓ 1730ms | ✓ 1853ms | ✓ 1752ms | http |
| 1.234.153.14:80 | ✓ 896ms | ✓ 1179ms | ✓ 908ms | ✓ 1164ms | ✓ 866ms | http |
| 8.137.112.117:3128 | ✓ 937ms | 否 | ✓ 1220ms | ✓ 1328ms | ✓ 1157ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1478ms | ✓ 1637ms | ✓ 1177ms | http |
| 103.82.23.118:5261 | ✓ 1909ms | 否 | ✓ 1558ms | 否 | ✓ 1906ms | http |
| 177.93.132.244:3128 | ✓ 737ms | 否 | ✓ 633ms | 否 | ✓ 1628ms | http |
| 168.144.75.9:3128 | ✓ 1804ms | 否 | ✓ 1717ms | ✓ 1792ms | 否 | http |
| 8.209.238.110:47701 | ✓ 652ms | ✓ 1304ms | ✓ 788ms | ✓ 1035ms | ✓ 863ms | http |
| 60.249.94.208:3128 | ✓ 1054ms | ✓ 1148ms | ✓ 916ms | ✓ 1174ms | ✓ 919ms | http |
| 187.248.75.106:8081 | ✓ 712ms | ✓ 1226ms | 否 | ✓ 1513ms | 否 | http |
| 146.19.56.212:40002 | ✓ 861ms | ✓ 1472ms | ✓ 754ms | 否 | 否 | http |
| 47.105.98.23:3128 | ✓ 873ms | ✓ 1129ms | ✓ 954ms | ✓ 1199ms | ✓ 928ms | http |
| 121.230.8.220:1080 | 否 | 否 | ✓ 1260ms | ✓ 1745ms | ✓ 1781ms | http |
| 108.181.0.167:8080 | ✓ 637ms | ✓ 1200ms | ✓ 336ms | ✓ 1166ms | ✓ 1580ms | http |
| 103.112.163.46:8080 | 否 | 否 | ✓ 1994ms | ✓ 1686ms | ✓ 1999ms | http |
| 130.61.139.145:3128 | ✓ 884ms | 否 | ✓ 1999ms | ✓ 1677ms | ✓ 1230ms | http |

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
