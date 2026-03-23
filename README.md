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

最后更新：2026-03-23 08:07:45 UTC（2026-03-23 16:07:45 UTC+8）

**代理总数：67**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 66 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 67 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1522ms | ✓ 1818ms | ✓ 789ms | ✓ 1161ms | ✓ 1119ms | http |
| 167.103.34.108:8800 | ✓ 1661ms | 否 | ✓ 1467ms | ✓ 1822ms | 否 | http |
| 45.167.124.52:8080 | ✓ 896ms | 否 | ✓ 1688ms | ✓ 1895ms | ✓ 1708ms | http |
| 35.225.22.61:80 | ✓ 880ms | 否 | ✓ 711ms | 否 | ✓ 785ms | http |
| 103.113.70.189:1081 | ✓ 335ms | 否 | ✓ 195ms | ✓ 1342ms | ✓ 936ms | http |
| 43.99.54.236:5555 | ✓ 806ms | ✓ 1107ms | ✓ 768ms | ✓ 978ms | ✓ 760ms | http |
| 43.153.28.68:3128 | ✓ 740ms | 否 | ✓ 538ms | ✓ 905ms | 否 | http |
| 113.160.132.26:8080 | ✓ 1513ms | 否 | ✓ 1438ms | ✓ 1686ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1935ms | 否 | ✓ 1788ms | ✓ 1981ms | 否 | http |
| 116.171.106.15:3443 | 否 | 否 | ✓ 1706ms | ✓ 1999ms | ✓ 1703ms | http |
| 115.231.181.40:8128 | ✓ 1659ms | 否 | ✓ 1090ms | 否 | ✓ 1400ms | http |
| 147.161.239.240:8800 | ✓ 1061ms | 否 | ✓ 786ms | ✓ 1698ms | ✓ 1344ms | http |
| 120.92.212.16:8890 | ✓ 1136ms | ✓ 1347ms | ✓ 1295ms | ✓ 1331ms | ✓ 1280ms | http |
| 203.194.114.217:3128 | ✓ 1636ms | 否 | ✓ 1613ms | ✓ 1362ms | ✓ 1184ms | http |
| 160.250.4.245:1 | ✓ 1646ms | 否 | ✓ 1793ms | ✓ 1442ms | ✓ 1121ms | http |
| 101.43.127.100:8877 | ✓ 943ms | ✓ 1930ms | 否 | 否 | ✓ 931ms | http |
| 218.89.134.230:3333 | 否 | ✓ 1675ms | ✓ 1706ms | ✓ 1741ms | ✓ 1380ms | http |
| 142.171.224.229:7890 | ✓ 341ms | ✓ 951ms | ✓ 1071ms | ✓ 842ms | ✓ 646ms | http |
| 38.34.179.61:8445 | 否 | ✓ 1677ms | ✓ 1287ms | ✓ 870ms | 否 | http |
| 38.34.179.96:8451 | 否 | ✓ 1287ms | ✓ 1195ms | ✓ 995ms | ✓ 1016ms | http |
| 77.110.113.24:40000 | ✓ 788ms | 否 | ✓ 1522ms | 否 | ✓ 1987ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1372ms | ✓ 1104ms | ✓ 1349ms | ✓ 1829ms | http |
| 219.117.204.211:7799 | 否 | ✓ 1610ms | ✓ 876ms | ✓ 1606ms | 否 | http |
| 1.231.81.166:3128 | ✓ 1029ms | ✓ 1293ms | ✓ 1377ms | ✓ 1476ms | ✓ 1475ms | http |
| 101.47.73.135:3128 | 否 | 否 | ✓ 1337ms | ✓ 1564ms | ✓ 1365ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1657ms | ✓ 1611ms | ✓ 1226ms | http |
| 38.145.218.228:8447 | 否 | ✓ 907ms | 否 | ✓ 1706ms | ✓ 1567ms | http |
| 64.227.76.27:1080 | ✓ 498ms | ✓ 1852ms | ✓ 1858ms | 否 | 否 | http |
| 172.212.68.37:3128 | ✓ 313ms | 否 | ✓ 1037ms | ✓ 1120ms | ✓ 1165ms | http |
| 38.34.179.150:8444 | ✓ 1494ms | ✓ 1322ms | 否 | ✓ 1439ms | ✓ 1567ms | http |
| 58.220.95.8:10174 | ✓ 961ms | ✓ 1317ms | 否 | ✓ 1600ms | ✓ 1026ms | http |
| 122.233.92.199:1111 | ✓ 975ms | ✓ 1313ms | ✓ 1421ms | 否 | ✓ 1854ms | http |
| 38.145.220.103:8452 | ✓ 1191ms | 否 | ✓ 594ms | 否 | ✓ 1732ms | http |
| 38.145.203.109:8450 | ✓ 1820ms | ✓ 1960ms | ✓ 594ms | 否 | ✓ 1325ms | http |
| 38.145.218.163:8448 | ✓ 1516ms | ✓ 1288ms | 否 | ✓ 1284ms | 否 | http |
| 150.249.255.91:3128 | ✓ 1463ms | ✓ 1644ms | 否 | ✓ 990ms | ✓ 802ms | http |
| 137.220.151.110:6005 | ✓ 1812ms | 否 | ✓ 937ms | ✓ 1316ms | ✓ 977ms | http |
| 137.220.150.104:6005 | 否 | 否 | ✓ 1015ms | ✓ 1306ms | ✓ 1194ms | http |
| 166.88.55.83:7890 | ✓ 1495ms | ✓ 1229ms | ✓ 763ms | ✓ 929ms | ✓ 734ms | http |
| 202.141.161.53:30001 | ✓ 1985ms | 否 | ✓ 1361ms | ✓ 1323ms | ✓ 1222ms | http |
| 38.180.135.112:3128 | ✓ 388ms | ✓ 1022ms | ✓ 1959ms | ✓ 955ms | ✓ 942ms | http |
| 38.34.179.57:8453 | ✓ 1101ms | 否 | ✓ 1039ms | 否 | ✓ 1846ms | http |
| 38.34.179.150:8449 | ✓ 1212ms | 否 | ✓ 1514ms | ✓ 901ms | ✓ 1239ms | http |
| 111.79.111.126:3128 | 否 | ✓ 1236ms | ✓ 1093ms | ✓ 1337ms | ✓ 1365ms | http |
| 62.113.119.14:8080 | ✓ 745ms | 否 | ✓ 967ms | ✓ 1530ms | ✓ 1835ms | http |
| 43.165.195.107:3128 | ✓ 1567ms | ✓ 1627ms | ✓ 1651ms | ✓ 1360ms | ✓ 1078ms | http |
| 47.74.226.8:5001 | ✓ 1302ms | 否 | ✓ 1170ms | ✓ 1680ms | 否 | http |
| 185.191.236.162:3128 | ✓ 606ms | 否 | ✓ 1472ms | ✓ 1656ms | ✓ 1052ms | http |
| 222.228.171.92:8080 | 否 | 否 | ✓ 863ms | ✓ 1008ms | ✓ 931ms | http |
| 38.34.179.75:8453 | 否 | ✓ 1533ms | 否 | ✓ 1075ms | ✓ 1099ms | http |
| 194.67.99.223:1080 | ✓ 924ms | ✓ 1614ms | ✓ 1804ms | 否 | 否 | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1132ms | ✓ 1590ms | ✓ 1291ms | http |
| 103.39.51.190:8080 | ✓ 1957ms | 否 | ✓ 1937ms | ✓ 1458ms | ✓ 1464ms | http |
| 38.34.179.20:8445 | 否 | ✓ 1147ms | ✓ 1643ms | 否 | ✓ 858ms | http |
| 20.120.225.109:3128 | ✓ 590ms | ✓ 1374ms | ✓ 916ms | ✓ 1268ms | ✓ 893ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 1778ms | ✓ 1495ms | ✓ 754ms | http |
| 59.46.216.131:30001 | ✓ 1821ms | ✓ 1508ms | ✓ 1448ms | 否 | 否 | http |
| 137.220.150.170:6005 | ✓ 924ms | 否 | ✓ 1102ms | 否 | ✓ 1445ms | http |
| 38.34.179.172:8451 | ✓ 322ms | 否 | ✓ 1582ms | ✓ 893ms | 否 | http |
| 185.114.73.2:1080 | ✓ 1130ms | 否 | ✓ 1558ms | ✓ 1626ms | ✓ 1148ms | http |
| 113.176.92.71:3128 | ✓ 1841ms | ✓ 1696ms | 否 | ✓ 1531ms | 否 | http |
| 200.26.232.82:3128 | 否 | 否 | ✓ 1428ms | ✓ 1762ms | ✓ 1934ms | http |
| 91.233.223.147:3128 | ✓ 1357ms | 否 | ✓ 1392ms | 否 | ✓ 1829ms | http |
| 181.78.44.63:999 | ✓ 887ms | ✓ 1334ms | ✓ 1292ms | ✓ 1520ms | ✓ 1409ms | http |
| 116.80.49.170:3172 | ✓ 1696ms | 否 | ✓ 1597ms | ✓ 1963ms | 否 | http |
| 88.80.150.82:8080 | ✓ 798ms | 否 | ✓ 769ms | ✓ 1575ms | ✓ 1398ms | https |
| 103.139.138.194:3128 | 否 | 否 | ✓ 1542ms | ✓ 1704ms | ✓ 1280ms | http |

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
