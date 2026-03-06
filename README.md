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

最后更新：2026-03-06 17:45:07 UTC（2026-03-07 01:45:07 UTC+8）

**代理总数：77**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 76 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 314ms | ✓ 1916ms | 否 | ✓ 1248ms | ✓ 1046ms | http |
| 1.231.81.166:3128 | ✓ 1620ms | ✓ 1810ms | ✓ 954ms | ✓ 1053ms | ✓ 771ms | http |
| 152.42.195.165:8888 | ✓ 797ms | 否 | ✓ 836ms | ✓ 1353ms | ✓ 1059ms | http |
| 23.94.182.50:12345 | ✓ 973ms | 否 | ✓ 1866ms | ✓ 1491ms | ✓ 1236ms | http |
| 35.225.22.61:80 | ✓ 706ms | 否 | 否 | ✓ 984ms | ✓ 727ms | http |
| 103.84.95.54:7890 | 否 | 否 | ✓ 1606ms | ✓ 1069ms | ✓ 814ms | http |
| 45.22.209.157:8888 | ✓ 1533ms | 否 | 否 | ✓ 1937ms | ✓ 1402ms | http |
| 186.148.180.46:999 | 否 | 否 | ✓ 1386ms | ✓ 1943ms | ✓ 1756ms | http |
| 46.183.25.8:443 | ✓ 830ms | 否 | ✓ 1155ms | ✓ 1970ms | 否 | http |
| 136.49.39.94:8888 | ✓ 635ms | 否 | ✓ 1403ms | ✓ 1457ms | 否 | http |
| 113.176.92.71:3128 | ✓ 1828ms | ✓ 1411ms | ✓ 1328ms | ✓ 1571ms | ✓ 1072ms | http |
| 162.248.165.72:1080 | ✓ 1512ms | 否 | ✓ 960ms | 否 | ✓ 1843ms | http |
| 121.128.121.54:3128 | ✓ 1551ms | 否 | ✓ 709ms | ✓ 1041ms | ✓ 856ms | http |
| 194.213.18.200:443 | ✓ 934ms | 否 | 否 | ✓ 1217ms | ✓ 1856ms | http |
| 159.223.42.219:3128 | ✓ 890ms | 否 | ✓ 1164ms | ✓ 1157ms | 否 | http |
| 67.169.98.211:443 | ✓ 959ms | 否 | ✓ 1967ms | ✓ 1798ms | 否 | http |
| 61.72.110.54:3128 | ✓ 1422ms | 否 | ✓ 1567ms | 否 | ✓ 1428ms | http |
| 14.56.107.244:3128 | ✓ 710ms | 否 | ✓ 1663ms | 否 | ✓ 876ms | http |
| 125.128.12.144:3128 | ✓ 1484ms | ✓ 1242ms | 否 | ✓ 1958ms | ✓ 1983ms | http |
| 42.96.16.158:1311 | ✓ 1580ms | 否 | ✓ 1876ms | ✓ 1276ms | 否 | http |
| 167.172.69.123:80 | ✓ 961ms | 否 | ✓ 1214ms | ✓ 1443ms | ✓ 1676ms | http |
| 167.172.69.123:8080 | ✓ 963ms | 否 | ✓ 1203ms | ✓ 1539ms | ✓ 1777ms | http |
| 91.193.240.157:9877 | ✓ 1109ms | 否 | ✓ 1588ms | 否 | ✓ 1707ms | http |
| 125.128.12.14:3128 | ✓ 1611ms | 否 | ✓ 1472ms | 否 | ✓ 1203ms | http |
| 120.92.212.16:7890 | ✓ 995ms | ✓ 1292ms | 否 | ✓ 1307ms | 否 | http |
| 103.104.99.29:80 | ✓ 1838ms | 否 | 否 | ✓ 1665ms | ✓ 1687ms | http |
| 103.104.99.89:80 | ✓ 1838ms | 否 | ✓ 1843ms | ✓ 1607ms | 否 | http |
| 190.12.150.244:999 | ✓ 1409ms | 否 | ✓ 1531ms | ✓ 1686ms | ✓ 1981ms | http |
| 168.235.110.63:3128 | ✓ 312ms | 否 | ✓ 1016ms | ✓ 1782ms | 否 | http |
| 14.56.177.44:3128 | ✓ 1619ms | 否 | ✓ 1166ms | ✓ 1114ms | ✓ 921ms | http |
| 185.191.236.162:3128 | ✓ 1765ms | 否 | ✓ 1159ms | ✓ 1685ms | ✓ 1164ms | http |
| 14.225.222.164:7890 | ✓ 1457ms | 否 | ✓ 1947ms | ✓ 1264ms | 否 | http |
| 193.108.118.190:8888 | ✓ 1497ms | 否 | ✓ 1165ms | 否 | ✓ 1934ms | http |
| 101.43.255.96:80 | 否 | 否 | ✓ 1389ms | ✓ 1679ms | ✓ 1017ms | http |
| 45.140.147.82:1081 | ✓ 542ms | 否 | ✓ 842ms | ✓ 1888ms | ✓ 1126ms | http |
| 124.156.179.148:3128 | ✓ 726ms | ✓ 986ms | ✓ 723ms | ✓ 903ms | ✓ 736ms | http |
| 202.155.12.161:443 | 否 | 否 | ✓ 1369ms | ✓ 1221ms | ✓ 1120ms | http |
| 61.72.110.94:3128 | ✓ 1335ms | 否 | ✓ 1015ms | ✓ 1601ms | 否 | http |
| 61.72.221.94:3128 | ✓ 1871ms | 否 | ✓ 1538ms | ✓ 1349ms | ✓ 1323ms | http |
| 107.174.80.186:3128 | ✓ 399ms | 否 | ✓ 387ms | ✓ 1015ms | ✓ 747ms | http |
| 5.9.55.221:5000 | ✓ 717ms | ✓ 1918ms | ✓ 1485ms | 否 | 否 | http |
| 5.252.33.13:2025 | ✓ 1463ms | 否 | ✓ 1496ms | 否 | ✓ 1824ms | http |
| 42.115.72.27:2064 | ✓ 1760ms | 否 | ✓ 1996ms | 否 | ✓ 1788ms | http |
| 120.92.212.16:8890 | 否 | 否 | ✓ 1089ms | ✓ 1545ms | ✓ 1284ms | http |
| 61.72.221.194:3128 | 否 | 否 | ✓ 1111ms | ✓ 1305ms | ✓ 1585ms | http |
| 138.124.53.25:7443 | ✓ 1828ms | 否 | ✓ 1855ms | 否 | ✓ 1724ms | http |
| 61.72.221.234:3128 | ✓ 1427ms | 否 | ✓ 1390ms | 否 | ✓ 876ms | http |
| 178.236.245.17:3128 | ✓ 1368ms | 否 | ✓ 1054ms | ✓ 1731ms | ✓ 1886ms | http |
| 178.236.245.59:3128 | ✓ 1368ms | 否 | ✓ 1056ms | ✓ 1747ms | ✓ 1895ms | http |
| 115.231.181.40:8128 | ✓ 1042ms | ✓ 1346ms | 否 | ✓ 1788ms | ✓ 1276ms | http |
| 42.115.72.27:2065 | ✓ 1682ms | 否 | ✓ 1630ms | ✓ 1908ms | ✓ 1552ms | http |
| 42.115.72.27:2102 | ✓ 1791ms | 否 | ✓ 1803ms | ✓ 1809ms | ✓ 1592ms | http |
| 2.56.178.131:443 | ✓ 1036ms | 否 | ✓ 1083ms | ✓ 1911ms | 否 | http |
| 103.139.138.194:3128 | ✓ 1977ms | 否 | ✓ 1962ms | 否 | ✓ 1823ms | http |
| 154.53.40.110:3128 | ✓ 668ms | 否 | ✓ 1791ms | ✓ 1146ms | ✓ 1143ms | http |
| 119.46.68.228:10227 | ✓ 1774ms | 否 | ✓ 1651ms | ✓ 1373ms | ✓ 1084ms | http |
| 119.46.68.239:10227 | ✓ 1771ms | 否 | 否 | ✓ 1375ms | ✓ 1095ms | http |
| 193.168.173.136:443 | ✓ 849ms | ✓ 1849ms | ✓ 1141ms | ✓ 1913ms | 否 | http |
| 210.223.44.230:3128 | 否 | ✓ 1856ms | ✓ 1073ms | ✓ 1689ms | ✓ 1939ms | http |
| 45.136.198.40:3128 | ✓ 1272ms | ✓ 1744ms | 否 | 否 | ✓ 1875ms | http |
| 42.115.72.27:2039 | 否 | 否 | ✓ 1654ms | ✓ 1825ms | ✓ 1628ms | http |
| 91.107.175.112:10801 | ✓ 497ms | 否 | ✓ 951ms | ✓ 1685ms | 否 | http |
| 119.46.68.228:443 | ✓ 1689ms | 否 | ✓ 1069ms | ✓ 1420ms | ✓ 1091ms | http |
| 81.70.169.194:80 | ✓ 1036ms | 否 | ✓ 1969ms | 否 | ✓ 1097ms | http |
| 59.46.216.131:30001 | ✓ 1081ms | ✓ 1739ms | 否 | ✓ 1547ms | 否 | http |
| 103.215.36.88:13763 | ✓ 1032ms | ✓ 1396ms | 否 | ✓ 1349ms | ✓ 1357ms | http |
| 172.212.68.37:3128 | ✓ 282ms | 否 | ✓ 1638ms | 否 | ✓ 1367ms | http |
| 119.46.68.228:80 | ✓ 1008ms | 否 | ✓ 975ms | ✓ 1371ms | ✓ 1084ms | http |
| 107.152.32.98:1305 | ✓ 1906ms | 否 | ✓ 1216ms | 否 | ✓ 1451ms | http |
| 103.215.36.88:18989 | 否 | 否 | ✓ 1053ms | ✓ 1497ms | ✓ 1110ms | http |
| 103.39.51.190:8080 | ✓ 1770ms | 否 | 否 | ✓ 1696ms | ✓ 1518ms | http |
| 45.140.147.155:1081 | ✓ 633ms | ✓ 1337ms | ✓ 1414ms | ✓ 1359ms | ✓ 1129ms | http |
| 150.249.255.91:3128 | ✓ 1580ms | 否 | ✓ 678ms | ✓ 1072ms | ✓ 781ms | http |
| 103.82.23.118:5235 | 否 | 否 | ✓ 1497ms | ✓ 1667ms | ✓ 1421ms | http |
| 61.52.131.172:8443 | ✓ 922ms | ✓ 1884ms | ✓ 1040ms | ✓ 1212ms | 否 | http |
| 103.82.23.118:5247 | ✓ 1606ms | 否 | ✓ 1541ms | ✓ 1553ms | ✓ 1532ms | http |
| 88.80.150.82:8080 | ✓ 831ms | 否 | ✓ 1123ms | ✓ 1648ms | 否 | https |

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
