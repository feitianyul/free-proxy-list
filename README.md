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

最后更新：2026-05-20 14:33:15 UTC（2026-05-20 22:33:15 UTC+8）

**代理总数：56**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 56 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 56 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 192.99.8.15:8850 | ✓ 941ms | 否 | ✓ 1063ms | ✓ 1301ms | ✓ 1056ms | http |
| 89.58.50.94:11140 | ✓ 1079ms | 否 | ✓ 1400ms | ✓ 1833ms | ✓ 1719ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1906ms | ✓ 1077ms | ✓ 1806ms | 否 | http |
| 170.106.136.181:31002 | ✓ 477ms | ✓ 1177ms | ✓ 410ms | ✓ 1761ms | ✓ 699ms | http |
| 176.111.37.216:39811 | ✓ 943ms | ✓ 1478ms | ✓ 1126ms | 否 | 否 | http |
| 176.111.37.5:39811 | ✓ 707ms | ✓ 1749ms | ✓ 661ms | 否 | ✓ 1472ms | http |
| 45.117.163.134:3128 | ✓ 1311ms | 否 | ✓ 1025ms | ✓ 1352ms | ✓ 1043ms | http |
| 174.137.134.182:2999 | ✓ 253ms | ✓ 1037ms | ✓ 116ms | ✓ 1905ms | ✓ 1739ms | http |
| 43.130.126.146:6688 | ✓ 1281ms | ✓ 843ms | ✓ 74ms | 否 | ✓ 881ms | http |
| 138.2.92.70:8100 | 否 | 否 | ✓ 1591ms | ✓ 1943ms | ✓ 1509ms | http |
| 185.200.188.234:10001 | ✓ 1671ms | 否 | ✓ 748ms | 否 | ✓ 1567ms | http |
| 138.2.78.251:8100 | 否 | 否 | ✓ 1935ms | ✓ 1887ms | ✓ 1709ms | http |
| 74.208.192.81:3129 | ✓ 1342ms | 否 | ✓ 348ms | ✓ 1061ms | 否 | http |
| 128.199.254.13:9090 | ✓ 1158ms | 否 | ✓ 1667ms | ✓ 1339ms | ✓ 1047ms | http |
| 128.199.113.85:9090 | ✓ 1163ms | 否 | ✓ 1264ms | ✓ 1348ms | ✓ 1105ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1866ms | ✓ 1800ms | ✓ 1589ms | ✓ 1226ms | http |
| 202.28.194.139:31280 | ✓ 1840ms | 否 | 否 | ✓ 1936ms | ✓ 1925ms | http |
| 38.250.126.225:999 | ✓ 1392ms | 否 | ✓ 783ms | 否 | ✓ 1837ms | http |
| 128.199.116.219:9090 | ✓ 984ms | 否 | ✓ 1496ms | ✓ 1334ms | ✓ 1036ms | http |
| 45.125.67.37:8443 | ✓ 1129ms | 否 | ✓ 1448ms | ✓ 1918ms | ✓ 1372ms | http |
| 148.230.4.241:999 | 否 | 否 | ✓ 635ms | ✓ 1369ms | ✓ 1279ms | http |
| 8.154.21.175:3128 | ✓ 1127ms | ✓ 1369ms | ✓ 1154ms | ✓ 1371ms | 否 | http |
| 34.87.80.221:30000 | ✓ 1449ms | 否 | ✓ 1539ms | ✓ 1308ms | ✓ 1105ms | http |
| 168.110.52.228:3128 | ✓ 1751ms | 否 | ✓ 1934ms | ✓ 1244ms | ✓ 1353ms | http |
| 114.214.165.78:10810 | ✓ 1375ms | 否 | ✓ 1506ms | ✓ 1793ms | ✓ 1413ms | http |
| 121.230.8.235:1080 | 否 | ✓ 1713ms | ✓ 1620ms | 否 | ✓ 1543ms | http |
| 34.101.184.164:3128 | 否 | 否 | ✓ 1758ms | ✓ 1548ms | ✓ 1233ms | http |
| 152.67.191.232:6800 | ✓ 1118ms | 否 | ✓ 1108ms | ✓ 1606ms | ✓ 1262ms | http |
| 128.199.114.189:9090 | ✓ 1170ms | 否 | ✓ 1515ms | ✓ 1403ms | ✓ 1069ms | http |
| 62.113.119.14:8080 | ✓ 1256ms | 否 | ✓ 901ms | ✓ 1463ms | 否 | http |
| 119.23.68.90:9003 | ✓ 1209ms | ✓ 1472ms | ✓ 1175ms | 否 | ✓ 1083ms | http |
| 185.125.231.40:31280 | ✓ 1007ms | 否 | ✓ 1654ms | 否 | ✓ 1767ms | http |
| 185.104.249.25:3128 | ✓ 876ms | 否 | ✓ 1853ms | 否 | ✓ 1719ms | http |
| 3.101.133.120:80 | ✓ 1047ms | ✓ 1632ms | ✓ 862ms | 否 | ✓ 1133ms | http |
| 38.19.41.140:999 | ✓ 1420ms | 否 | ✓ 1492ms | ✓ 1920ms | ✓ 1349ms | http |
| 38.19.41.141:999 | ✓ 1333ms | ✓ 1544ms | 否 | ✓ 1946ms | ✓ 1668ms | http |
| 38.19.41.126:999 | ✓ 1334ms | ✓ 1716ms | 否 | ✓ 1847ms | ✓ 1640ms | http |
| 103.157.79.155:1818 | ✓ 1986ms | 否 | ✓ 1551ms | ✓ 1726ms | ✓ 1696ms | http |
| 38.19.41.121:999 | ✓ 1680ms | 否 | ✓ 1820ms | 否 | ✓ 1765ms | http |
| 38.19.41.120:999 | ✓ 1669ms | 否 | ✓ 1838ms | ✓ 1852ms | ✓ 1981ms | http |
| 20.210.39.153:8561 | 否 | 否 | ✓ 1627ms | ✓ 1778ms | ✓ 1369ms | http |
| 114.214.163.108:6789 | ✓ 1294ms | ✓ 1660ms | ✓ 1559ms | ✓ 1630ms | ✓ 1286ms | http |
| 45.173.12.140:1994 | ✓ 1328ms | 否 | 否 | ✓ 1737ms | ✓ 1625ms | http |
| 77.110.107.80:1080 | ✓ 1174ms | 否 | ✓ 1911ms | ✓ 1540ms | ✓ 1115ms | http |
| 77.110.107.80:8080 | ✓ 1165ms | 否 | ✓ 1911ms | ✓ 1578ms | ✓ 1077ms | http |
| 147.45.78.89:1080 | ✓ 1182ms | 否 | ✓ 1199ms | 否 | ✓ 1721ms | http |
| 173.212.245.136:8888 | ✓ 1638ms | 否 | ✓ 1793ms | 否 | ✓ 1852ms | http |
| 152.32.132.190:7890 | ✓ 1265ms | 否 | ✓ 1768ms | 否 | ✓ 1807ms | http |
| 158.160.215.167:8127 | ✓ 1065ms | 否 | ✓ 1561ms | 否 | ✓ 1639ms | http |
| 152.42.170.187:9090 | 否 | 否 | ✓ 1641ms | ✓ 1416ms | ✓ 1016ms | http |
| 104.248.151.93:9090 | ✓ 1389ms | 否 | ✓ 1924ms | 否 | ✓ 1449ms | http |
| 103.82.23.118:5347 | 否 | 否 | ✓ 1981ms | ✓ 1849ms | ✓ 1809ms | http |
| 128.199.121.61:9090 | ✓ 1208ms | 否 | ✓ 1341ms | ✓ 1372ms | 否 | http |
| 105.159.148.165:5779 | ✓ 1178ms | ✓ 1658ms | ✓ 1751ms | ✓ 1484ms | ✓ 1400ms | http |
| 61.52.131.172:8443 | ✓ 1087ms | 否 | ✓ 1722ms | ✓ 1437ms | ✓ 1171ms | http |
| 57.128.188.167:8181 | ✓ 1689ms | 否 | ✓ 1389ms | ✓ 1853ms | 否 | http |

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
