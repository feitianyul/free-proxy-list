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

最后更新：2026-03-09 05:53:42 UTC（2026-03-09 13:53:42 UTC+8）

**代理总数：57**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 56 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 57 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1895ms | ✓ 1482ms | ✓ 1126ms | ✓ 1456ms | ✓ 981ms | http |
| 101.47.73.135:3128 | ✓ 1693ms | 否 | ✓ 1443ms | ✓ 1517ms | 否 | http |
| 85.208.108.43:2094 | ✓ 1402ms | 否 | ✓ 1260ms | ✓ 1171ms | ✓ 738ms | http |
| 85.208.108.43:10808 | ✓ 1357ms | 否 | ✓ 1436ms | ✓ 993ms | ✓ 1736ms | http |
| 165.227.5.10:8888 | ✓ 1650ms | ✓ 1074ms | ✓ 1441ms | ✓ 902ms | ✓ 839ms | http |
| 152.42.213.210:8080 | ✓ 1474ms | 否 | ✓ 1330ms | ✓ 1359ms | ✓ 1274ms | http |
| 121.126.185.63:25152 | 否 | 否 | ✓ 1363ms | ✓ 1989ms | ✓ 1516ms | http |
| 120.92.212.16:8890 | ✓ 1159ms | 否 | 否 | ✓ 1361ms | ✓ 1137ms | http |
| 202.155.12.161:443 | 否 | 否 | ✓ 1309ms | ✓ 1070ms | ✓ 941ms | http |
| 194.213.18.200:443 | ✓ 1627ms | 否 | 否 | ✓ 1053ms | ✓ 878ms | http |
| 81.70.169.194:80 | ✓ 1208ms | ✓ 1449ms | ✓ 1177ms | 否 | 否 | http |
| 5.129.206.247:8888 | ✓ 1484ms | 否 | ✓ 1640ms | 否 | ✓ 1940ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1095ms | ✓ 1403ms | ✓ 1083ms | http |
| 101.43.255.96:80 | ✓ 1086ms | 否 | ✓ 1577ms | 否 | ✓ 1857ms | http |
| 103.84.95.54:7890 | ✓ 1511ms | 否 | ✓ 1846ms | ✓ 1338ms | ✓ 881ms | http |
| 62.113.119.14:8080 | ✓ 601ms | 否 | ✓ 595ms | ✓ 1487ms | 否 | http |
| 46.249.103.192:443 | ✓ 548ms | 否 | ✓ 799ms | 否 | ✓ 1567ms | http |
| 160.250.5.22:1 | ✓ 1744ms | 否 | ✓ 1558ms | ✓ 1702ms | ✓ 1173ms | http |
| 160.250.4.245:1 | ✓ 1721ms | 否 | ✓ 1545ms | ✓ 1479ms | ✓ 1472ms | http |
| 116.80.82.218:3172 | ✓ 1730ms | 否 | ✓ 1639ms | ✓ 1983ms | 否 | http |
| 86.109.3.24:10012 | ✓ 760ms | 否 | ✓ 1291ms | ✓ 1198ms | 否 | http |
| 121.230.9.26:1080 | ✓ 1201ms | 否 | ✓ 1298ms | ✓ 1451ms | 否 | http |
| 115.231.181.40:8128 | ✓ 1312ms | 否 | ✓ 1384ms | 否 | ✓ 1980ms | http |
| 168.235.110.63:3128 | ✓ 946ms | 否 | ✓ 1164ms | ✓ 1048ms | ✓ 922ms | http |
| 190.9.109.207:999 | 否 | 否 | ✓ 1160ms | ✓ 1471ms | ✓ 1081ms | http |
| 35.225.22.61:80 | 否 | ✓ 1595ms | ✓ 426ms | ✓ 977ms | ✓ 776ms | http |
| 45.136.198.40:3128 | ✓ 1258ms | 否 | ✓ 1534ms | ✓ 1912ms | ✓ 1428ms | http |
| 121.230.9.148:1080 | ✓ 1399ms | ✓ 1933ms | ✓ 1187ms | 否 | 否 | http |
| 121.230.8.171:1080 | 否 | ✓ 1771ms | ✓ 1181ms | 否 | ✓ 1579ms | http |
| 118.113.246.172:1080 | ✓ 1437ms | ✓ 1913ms | ✓ 1885ms | 否 | ✓ 1871ms | http |
| 159.223.42.219:3128 | 否 | 否 | ✓ 1360ms | ✓ 1389ms | ✓ 943ms | http |
| 91.233.223.147:3128 | ✓ 1652ms | 否 | ✓ 1249ms | 否 | ✓ 1731ms | http |
| 88.80.150.82:8080 | ✓ 1387ms | ✓ 1996ms | 否 | 否 | ✓ 1710ms | https |
| 103.82.23.118:5185 | ✓ 1540ms | 否 | ✓ 1277ms | ✓ 1587ms | ✓ 1390ms | http |
| 1.225.116.115:1080 | ✓ 979ms | ✓ 1454ms | ✓ 1189ms | ✓ 1858ms | ✓ 1271ms | http |
| 156.225.70.152:39151 | 否 | 否 | ✓ 1863ms | ✓ 1647ms | ✓ 1121ms | http |
| 222.228.171.92:8080 | ✓ 1990ms | 否 | ✓ 1653ms | 否 | ✓ 1006ms | http |
| 162.248.165.72:1080 | ✓ 1083ms | ✓ 1502ms | ✓ 1284ms | 否 | 否 | http |
| 152.42.213.210:80 | ✓ 1495ms | 否 | ✓ 1571ms | ✓ 1480ms | ✓ 1324ms | http |
| 45.129.141.143:3128 | ✓ 960ms | 否 | ✓ 1793ms | ✓ 1792ms | ✓ 1523ms | http |
| 120.238.159.234:22222 | 否 | ✓ 1413ms | ✓ 1013ms | ✓ 1320ms | ✓ 971ms | http |
| 209.38.51.97:3128 | ✓ 943ms | 否 | ✓ 902ms | 否 | ✓ 955ms | http |
| 20.120.225.109:3128 | ✓ 1144ms | ✓ 1291ms | ✓ 1836ms | 否 | ✓ 1059ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1961ms | ✓ 1409ms | ✓ 1489ms | http |
| 116.80.82.217:3172 | ✓ 1767ms | 否 | ✓ 1659ms | 否 | ✓ 1745ms | http |
| 150.107.140.238:3128 | ✓ 1696ms | 否 | ✓ 1156ms | 否 | ✓ 1151ms | http |
| 148.135.116.20:8118 | 否 | ✓ 1537ms | ✓ 1006ms | 否 | ✓ 945ms | http |
| 89.185.85.138:1080 | ✓ 442ms | 否 | ✓ 1965ms | ✓ 1502ms | ✓ 1939ms | http |
| 45.22.209.157:8888 | ✓ 795ms | 否 | ✓ 1443ms | 否 | ✓ 879ms | http |
| 47.101.149.27:9010 | ✓ 1907ms | 否 | ✓ 1941ms | ✓ 1664ms | 否 | http |
| 113.132.112.110:9000 | 否 | ✓ 1721ms | ✓ 1683ms | 否 | ✓ 1745ms | http |
| 47.77.193.180:1080 | ✓ 393ms | ✓ 982ms | ✓ 507ms | ✓ 899ms | ✓ 1275ms | http |
| 210.48.154.94:80 | ✓ 1812ms | 否 | ✓ 1535ms | ✓ 1441ms | ✓ 1436ms | http |
| 5.101.0.233:3128 | ✓ 1049ms | ✓ 1848ms | ✓ 1376ms | 否 | ✓ 1881ms | http |
| 106.14.203.63:3333 | 否 | ✓ 1886ms | ✓ 1105ms | 否 | ✓ 1767ms | http |
| 197.164.101.13:1976 | ✓ 1267ms | 否 | ✓ 1509ms | 否 | ✓ 1917ms | http |
| 138.197.68.35:4857 | ✓ 961ms | ✓ 1864ms | 否 | 否 | ✓ 1598ms | http |

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
