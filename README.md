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

最后更新：2026-03-29 09:43:52 UTC（2026-03-29 17:43:52 UTC+8）

**代理总数：50**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 49 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 50 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.99.54.236:5555 | ✓ 741ms | 否 | ✓ 748ms | ✓ 919ms | ✓ 729ms | http |
| 147.161.210.140:8800 | ✓ 1652ms | 否 | ✓ 863ms | ✓ 1033ms | ✓ 1055ms | http |
| 167.103.115.102:8800 | ✓ 1483ms | 否 | ✓ 1420ms | ✓ 1297ms | ✓ 1279ms | http |
| 167.103.34.108:8800 | ✓ 1757ms | 否 | ✓ 1597ms | ✓ 1686ms | 否 | http |
| 95.213.217.168:52004 | ✓ 627ms | ✓ 1668ms | ✓ 608ms | ✓ 1659ms | ✓ 1260ms | http |
| 113.160.132.26:8080 | ✓ 1652ms | 否 | ✓ 962ms | ✓ 1233ms | ✓ 1296ms | http |
| 45.140.147.155:1081 | ✓ 822ms | 否 | 否 | ✓ 1281ms | ✓ 1095ms | http |
| 167.103.144.127:8800 | ✓ 1334ms | 否 | ✓ 1638ms | ✓ 1723ms | 否 | http |
| 180.250.219.58:53281 | 否 | 否 | ✓ 1543ms | ✓ 1964ms | ✓ 1927ms | http |
| 167.103.31.122:8800 | ✓ 1514ms | 否 | ✓ 1291ms | ✓ 1690ms | 否 | http |
| 45.12.151.226:2829 | ✓ 1127ms | 否 | 否 | ✓ 1851ms | ✓ 1418ms | http |
| 120.92.212.16:8890 | ✓ 1117ms | ✓ 1307ms | ✓ 1016ms | 否 | 否 | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 1068ms | ✓ 966ms | ✓ 818ms | http |
| 147.161.239.240:8800 | ✓ 635ms | ✓ 1646ms | ✓ 1386ms | ✓ 1750ms | ✓ 1443ms | http |
| 1.231.81.166:3128 | ✓ 1688ms | ✓ 1055ms | ✓ 1635ms | ✓ 1139ms | ✓ 1053ms | http |
| 42.96.16.158:1311 | ✓ 1955ms | 否 | ✓ 1321ms | ✓ 1287ms | 否 | http |
| 115.231.181.40:8128 | 否 | ✓ 1199ms | ✓ 1002ms | ✓ 1352ms | 否 | http |
| 101.43.127.100:8877 | ✓ 985ms | ✓ 1174ms | ✓ 944ms | ✓ 1263ms | 否 | http |
| 34.101.184.164:3128 | ✓ 1790ms | 否 | ✓ 1565ms | ✓ 1817ms | ✓ 1885ms | http |
| 45.140.147.155:1082 | ✓ 828ms | ✓ 1323ms | ✓ 1541ms | ✓ 1734ms | ✓ 1206ms | http |
| 103.84.95.54:7890 | ✓ 877ms | 否 | ✓ 1285ms | ✓ 1034ms | ✓ 914ms | http |
| 103.82.23.118:5247 | ✓ 1534ms | 否 | ✓ 1124ms | ✓ 1732ms | ✓ 1665ms | http |
| 219.117.204.211:7799 | ✓ 1594ms | ✓ 1932ms | ✓ 1332ms | 否 | 否 | http |
| 120.92.212.16:7890 | ✓ 1971ms | ✓ 1289ms | ✓ 1121ms | 否 | 否 | http |
| 5.104.87.17:8051 | 否 | 否 | ✓ 1317ms | ✓ 1773ms | ✓ 1431ms | http |
| 59.46.216.131:30001 | ✓ 992ms | ✓ 1403ms | ✓ 1165ms | 否 | 否 | http |
| 208.87.243.199:7878 | ✓ 593ms | ✓ 1084ms | ✓ 1106ms | ✓ 1377ms | 否 | http |
| 38.145.220.33:8448 | ✓ 752ms | ✓ 1614ms | 否 | ✓ 921ms | ✓ 691ms | http |
| 121.230.9.205:1080 | ✓ 1076ms | ✓ 1599ms | ✓ 1123ms | 否 | ✓ 1382ms | http |
| 8.219.97.248:80 | ✓ 1997ms | 否 | 否 | ✓ 1804ms | ✓ 1450ms | http |
| 45.136.198.40:3128 | ✓ 1354ms | ✓ 1802ms | 否 | ✓ 1885ms | ✓ 1714ms | http |
| 177.234.217.88:999 | ✓ 1264ms | 否 | 否 | ✓ 1950ms | ✓ 1817ms | http |
| 38.145.208.191:8446 | ✓ 1267ms | ✓ 794ms | ✓ 619ms | ✓ 1470ms | ✓ 1276ms | http |
| 38.145.220.32:8450 | ✓ 377ms | ✓ 1621ms | ✓ 990ms | ✓ 1515ms | ✓ 655ms | http |
| 38.34.179.186:8453 | ✓ 1294ms | ✓ 1810ms | ✓ 1165ms | ✓ 1391ms | ✓ 1073ms | http |
| 38.34.179.20:8452 | ✓ 1040ms | ✓ 1256ms | 否 | ✓ 1675ms | 否 | http |
| 103.113.70.189:1081 | ✓ 454ms | 否 | ✓ 216ms | ✓ 1499ms | 否 | http |
| 38.145.220.41:8444 | ✓ 1259ms | ✓ 799ms | ✓ 1067ms | ✓ 1101ms | ✓ 1011ms | http |
| 64.227.76.27:1080 | ✓ 599ms | ✓ 1476ms | 否 | 否 | ✓ 1762ms | http |
| 194.67.99.223:1080 | ✓ 659ms | 否 | 否 | ✓ 1726ms | ✓ 1540ms | http |
| 103.39.51.190:8080 | ✓ 1877ms | 否 | 否 | ✓ 1440ms | ✓ 1811ms | http |
| 62.113.119.14:8080 | ✓ 1110ms | 否 | ✓ 767ms | ✓ 1633ms | ✓ 1228ms | http |
| 180.125.216.109:8118 | ✓ 980ms | 否 | ✓ 1116ms | 否 | ✓ 1010ms | http |
| 38.34.179.190:8451 | ✓ 392ms | ✓ 1688ms | ✓ 1648ms | ✓ 1401ms | ✓ 1645ms | http |
| 120.132.97.88:7897 | 否 | ✓ 1256ms | ✓ 994ms | ✓ 1206ms | ✓ 967ms | http |
| 89.208.106.138:10808 | ✓ 784ms | 否 | ✓ 1739ms | ✓ 1846ms | 否 | http |
| 113.176.92.71:3128 | ✓ 1516ms | ✓ 1440ms | ✓ 1281ms | ✓ 1903ms | ✓ 1460ms | http |
| 193.233.22.29:10808 | ✓ 508ms | 否 | ✓ 435ms | ✓ 1619ms | 否 | http |
| 114.237.77.231:1080 | ✓ 960ms | ✓ 1947ms | ✓ 1034ms | ✓ 1543ms | ✓ 1970ms | http |
| 88.80.150.82:8080 | ✓ 1700ms | ✓ 1912ms | 否 | 否 | ✓ 1779ms | https |

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
