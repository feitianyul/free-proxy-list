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

最后更新：2026-06-07 18:54:02 UTC（2026-06-08 02:54:02 UTC+8）

**代理总数：72**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 34.43.46.91:80 | ✓ 353ms | ✓ 1230ms | ✓ 1081ms | ✓ 1302ms | ✓ 1022ms | http |
| 185.200.188.234:10001 | ✓ 997ms | 否 | ✓ 702ms | 否 | ✓ 1603ms | http |
| 94.241.175.40:10808 | ✓ 959ms | 否 | ✓ 1765ms | 否 | ✓ 1265ms | http |
| 176.111.37.216:39811 | ✓ 1260ms | ✓ 1485ms | ✓ 933ms | 否 | 否 | http |
| 185.233.186.88:443 | ✓ 1222ms | 否 | ✓ 1623ms | 否 | ✓ 1591ms | http |
| 113.160.132.26:8080 | ✓ 1846ms | ✓ 1965ms | ✓ 1114ms | ✓ 1641ms | ✓ 1123ms | http |
| 38.123.220.147:999 | ✓ 790ms | ✓ 1269ms | ✓ 1061ms | ✓ 1879ms | ✓ 1290ms | http |
| 104.154.186.48:80 | ✓ 180ms | ✓ 1137ms | ✓ 819ms | ✓ 918ms | ✓ 960ms | http |
| 129.213.162.27:17777 | ✓ 402ms | ✓ 1677ms | 否 | ✓ 1138ms | ✓ 1002ms | http |
| 1.231.81.166:3128 | ✓ 875ms | ✓ 1149ms | ✓ 816ms | ✓ 1297ms | ✓ 1025ms | http |
| 185.141.26.131:3128 | ✓ 858ms | 否 | ✓ 1345ms | ✓ 1751ms | ✓ 1596ms | http |
| 81.200.154.236:48503 | ✓ 1074ms | 否 | 否 | ✓ 1658ms | ✓ 1187ms | http |
| 207.211.161.235:8888 | ✓ 499ms | ✓ 1672ms | 否 | ✓ 1816ms | 否 | http |
| 209.38.200.247:1080 | ✓ 464ms | ✓ 1979ms | ✓ 1203ms | ✓ 1534ms | ✓ 1304ms | http |
| 129.153.7.7:60000 | 否 | ✓ 868ms | 否 | ✓ 890ms | ✓ 1683ms | http |
| 50.114.102.16:8888 | ✓ 528ms | 否 | ✓ 405ms | ✓ 1457ms | ✓ 1073ms | http |
| 136.0.3.35:1234 | ✓ 576ms | 否 | ✓ 497ms | ✓ 1087ms | ✓ 1699ms | http |
| 85.234.100.149:8080 | ✓ 520ms | 否 | ✓ 885ms | 否 | ✓ 1158ms | http |
| 45.88.174.195:8080 | ✓ 1231ms | 否 | ✓ 1055ms | 否 | ✓ 1375ms | http |
| 43.128.145.26:1080 | 否 | 否 | ✓ 747ms | ✓ 1379ms | ✓ 1155ms | http |
| 202.28.194.139:31280 | ✓ 1813ms | 否 | ✓ 1833ms | ✓ 1828ms | ✓ 1929ms | http |
| 95.3.69.222:8080 | ✓ 1178ms | ✓ 1731ms | ✓ 994ms | ✓ 1775ms | ✓ 1460ms | http |
| 169.212.15.161:5000 | ✓ 1591ms | ✓ 1683ms | ✓ 1715ms | 否 | 否 | http |
| 14.143.222.113:10158 | ✓ 1757ms | 否 | ✓ 1004ms | ✓ 1446ms | 否 | http |
| 58.187.104.56:2104 | ✓ 1816ms | 否 | ✓ 1910ms | 否 | ✓ 1971ms | http |
| 165.227.133.230:8888 | ✓ 909ms | ✓ 1540ms | ✓ 827ms | 否 | 否 | http |
| 8.154.21.175:3128 | ✓ 1043ms | ✓ 1252ms | ✓ 1053ms | ✓ 1301ms | ✓ 1114ms | http |
| 37.49.224.15:3128 | ✓ 1000ms | 否 | ✓ 1536ms | ✓ 1783ms | 否 | http |
| 188.225.58.59:443 | ✓ 1477ms | 否 | ✓ 1193ms | 否 | ✓ 1373ms | http |
| 216.9.225.157:3128 | ✓ 314ms | 否 | ✓ 91ms | 否 | ✓ 1748ms | http |
| 147.45.78.89:1080 | 否 | 否 | ✓ 1974ms | ✓ 1659ms | ✓ 1251ms | http |
| 2.26.68.177:8080 | ✓ 499ms | ✓ 1916ms | ✓ 564ms | 否 | 否 | http |
| 170.106.136.181:31002 | ✓ 1690ms | ✓ 1018ms | ✓ 497ms | ✓ 906ms | ✓ 656ms | http |
| 39.100.88.235:3256 | 否 | ✓ 1450ms | ✓ 1144ms | ✓ 1475ms | ✓ 1164ms | http |
| 92.118.112.32:1082 | ✓ 955ms | ✓ 1823ms | ✓ 511ms | 否 | 否 | http |
| 116.104.252.1:2070 | 否 | 否 | ✓ 1585ms | ✓ 1822ms | ✓ 1782ms | http |
| 209.141.62.12:5555 | ✓ 503ms | ✓ 1401ms | 否 | ✓ 1135ms | 否 | http |
| 84.47.150.125:1080 | ✓ 1301ms | 否 | 否 | ✓ 1978ms | ✓ 1330ms | http |
| 62.210.136.222:3128 | ✓ 808ms | 否 | ✓ 853ms | ✓ 1904ms | ✓ 1507ms | http |
| 91.149.222.102:22335 | ✓ 1699ms | ✓ 1258ms | ✓ 1217ms | ✓ 1870ms | ✓ 1503ms | http |
| 185.103.103.156:1080 | ✓ 1866ms | 否 | ✓ 1363ms | 否 | ✓ 1901ms | http |
| 85.234.100.149:1080 | ✓ 1011ms | 否 | ✓ 1733ms | ✓ 1518ms | ✓ 1218ms | http |
| 138.124.93.170:3129 | 否 | 否 | ✓ 1554ms | ✓ 1762ms | ✓ 1527ms | http |
| 152.32.132.190:7890 | 否 | ✓ 1661ms | ✓ 1528ms | ✓ 1289ms | ✓ 1344ms | http |
| 46.8.112.212:3128 | ✓ 784ms | ✓ 1675ms | 否 | ✓ 1710ms | 否 | http |
| 92.118.112.32:1081 | 否 | ✓ 895ms | ✓ 902ms | ✓ 1114ms | ✓ 957ms | http |
| 103.157.117.226:81 | ✓ 1458ms | 否 | 否 | ✓ 1652ms | ✓ 1610ms | http |
| 43.134.141.85:80 | ✓ 1174ms | 否 | 否 | ✓ 1485ms | ✓ 1523ms | http |
| 80.150.246.98:443 | ✓ 502ms | ✓ 1440ms | ✓ 1411ms | 否 | ✓ 1553ms | http |
| 18.180.59.181:80 | ✓ 879ms | ✓ 1083ms | ✓ 1359ms | ✓ 1140ms | ✓ 1884ms | http |
| 91.233.223.147:3128 | ✓ 726ms | ✓ 1975ms | ✓ 718ms | ✓ 1888ms | ✓ 1440ms | http |
| es-xh-01.hpdata.click:443 | ✓ 1154ms | ✓ 1995ms | ✓ 1273ms | 否 | ✓ 1934ms | http |
| ch-xh-01.hpdata.click:443 | ✓ 1457ms | 否 | ✓ 1742ms | ✓ 1901ms | ✓ 1592ms | http |
| 183.80.40.243:2053 | 否 | 否 | ✓ 1694ms | ✓ 1960ms | ✓ 1765ms | http |
| 104.161.37.187:3128 | ✓ 886ms | 否 | ✓ 1282ms | 否 | ✓ 1438ms | http |
| 117.55.203.162:8899 | ✓ 1429ms | 否 | ✓ 1216ms | ✓ 1581ms | ✓ 1208ms | http |
| 117.55.203.161:8899 | ✓ 1434ms | 否 | ✓ 1199ms | ✓ 1828ms | ✓ 1240ms | http |
| 45.186.6.104:3128 | ✓ 1443ms | ✓ 1792ms | ✓ 1905ms | 否 | 否 | http |
| 116.104.252.1:2030 | ✓ 1604ms | 否 | 否 | ✓ 1834ms | ✓ 1624ms | http |
| 43.161.239.147:8888 | ✓ 817ms | 否 | ✓ 1690ms | ✓ 1705ms | 否 | http |
| 205.215.247.164:3128 | ✓ 1316ms | 否 | ✓ 739ms | ✓ 1167ms | ✓ 1458ms | http |
| 43.161.239.147:11090 | ✓ 1647ms | ✓ 1742ms | ✓ 1904ms | ✓ 1008ms | ✓ 1082ms | http |
| 3.90.0.161:8000 | ✓ 1361ms | 否 | ✓ 760ms | ✓ 1440ms | 否 | http |
| 171.238.103.148:2052 | ✓ 1577ms | 否 | ✓ 1818ms | 否 | ✓ 1633ms | http |
| 157.245.143.65:7890 | 否 | ✓ 949ms | 否 | ✓ 1854ms | ✓ 1705ms | http |
| 212.58.132.5:8888 | 否 | ✓ 1996ms | ✓ 1633ms | ✓ 1579ms | 否 | http |
| 43.228.215.32:8080 | ✓ 1104ms | 否 | ✓ 1186ms | ✓ 1319ms | ✓ 1188ms | http |
| 116.104.252.1:2087 | ✓ 1657ms | 否 | ✓ 1540ms | ✓ 1841ms | ✓ 1651ms | http |
| 61.52.131.172:8443 | ✓ 1023ms | ✓ 1382ms | ✓ 1170ms | ✓ 1359ms | ✓ 1106ms | http |
| 103.245.159.13:8080 | 否 | 否 | ✓ 1766ms | ✓ 1772ms | ✓ 1715ms | http |
| 103.209.36.58:8080 | ✓ 1578ms | 否 | 否 | ✓ 1745ms | ✓ 1864ms | http |
| 138.124.93.170:1080 | ✓ 852ms | ✓ 1402ms | ✓ 1121ms | ✓ 1750ms | ✓ 1455ms | http |

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
