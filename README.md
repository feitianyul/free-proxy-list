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

最后更新：2026-04-28 18:38:18 UTC（2026-04-29 02:38:18 UTC+8）

**代理总数：42**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 42 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 42 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 8.211.166.184:8081 | ✓ 1548ms | 否 | ✓ 916ms | ✓ 1117ms | ✓ 881ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1767ms | ✓ 1662ms | ✓ 1487ms | ✓ 1563ms | http |
| 45.167.124.71:999 | 否 | ✓ 1814ms | ✓ 1226ms | ✓ 1633ms | ✓ 1564ms | http |
| 115.231.181.40:8128 | ✓ 1072ms | ✓ 1375ms | 否 | ✓ 1508ms | ✓ 1144ms | http |
| 218.108.131.186:17890 | ✓ 1529ms | ✓ 1987ms | ✓ 1048ms | ✓ 1790ms | ✓ 1110ms | http |
| 8.154.21.175:3128 | ✓ 1530ms | 否 | ✓ 1074ms | ✓ 1315ms | ✓ 1075ms | http |
| 43.133.90.161:8888 | ✓ 1619ms | 否 | ✓ 1205ms | ✓ 1802ms | 否 | http |
| 217.76.245.80:999 | ✓ 804ms | ✓ 1537ms | ✓ 1056ms | ✓ 1343ms | ✓ 1043ms | http |
| 45.59.122.132:80 | ✓ 1609ms | 否 | ✓ 1823ms | 否 | ✓ 1384ms | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1221ms | ✓ 1620ms | ✓ 1289ms | http |
| 34.96.238.40:8080 | ✓ 1385ms | ✓ 1517ms | 否 | 否 | ✓ 1263ms | http |
| 160.250.4.245:1 | ✓ 1942ms | 否 | ✓ 1603ms | 否 | ✓ 1162ms | http |
| 47.85.51.197:1080 | ✓ 60ms | 否 | ✓ 303ms | ✓ 1212ms | ✓ 852ms | http |
| 46.101.95.183:8888 | ✓ 850ms | ✓ 1965ms | ✓ 560ms | 否 | 否 | http |
| 47.95.231.180:8084 | 否 | ✓ 1459ms | ✓ 1092ms | ✓ 1449ms | ✓ 1160ms | http |
| 114.55.226.123:10086 | ✓ 1237ms | ✓ 1595ms | ✓ 1356ms | ✓ 1643ms | ✓ 1359ms | http |
| 103.157.200.126:3128 | ✓ 1237ms | 否 | ✓ 1419ms | 否 | ✓ 1632ms | http |
| 45.140.147.155:1082 | ✓ 1746ms | 否 | ✓ 1481ms | 否 | ✓ 1391ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1863ms | ✓ 1718ms | ✓ 1816ms | http |
| 3.18.226.115:80 | 否 | ✓ 1786ms | ✓ 910ms | ✓ 1354ms | ✓ 1294ms | http |
| 101.32.244.83:8080 | ✓ 1157ms | 否 | ✓ 1137ms | ✓ 1745ms | ✓ 1487ms | http |
| 121.43.196.210:8222 | 否 | ✓ 1417ms | 否 | ✓ 1493ms | ✓ 1238ms | http |
| 62.113.119.14:8080 | ✓ 1259ms | ✓ 1623ms | ✓ 1266ms | ✓ 1461ms | ✓ 1131ms | http |
| 121.130.177.28:8888 | ✓ 1934ms | 否 | ✓ 1838ms | ✓ 1891ms | ✓ 1726ms | http |
| 121.130.199.80:3128 | ✓ 1384ms | ✓ 1223ms | ✓ 1447ms | ✓ 1384ms | ✓ 1362ms | http |
| 3.18.226.115:443 | ✓ 467ms | ✓ 1386ms | ✓ 1542ms | ✓ 1219ms | 否 | http |
| 195.133.28.208:3128 | ✓ 960ms | ✓ 1935ms | ✓ 784ms | ✓ 1480ms | ✓ 1164ms | http |
| 118.113.246.132:1080 | ✓ 1558ms | ✓ 1836ms | ✓ 1476ms | ✓ 1699ms | ✓ 1495ms | http |
| 107.173.122.179:10808 | ✓ 579ms | ✓ 1246ms | ✓ 1572ms | ✓ 1345ms | ✓ 879ms | http |
| 45.140.147.82:1081 | ✓ 525ms | ✓ 1420ms | ✓ 1627ms | ✓ 1368ms | ✓ 1625ms | http |
| 45.140.147.82:1082 | ✓ 485ms | ✓ 1388ms | ✓ 1699ms | ✓ 1380ms | 否 | http |
| 86.104.72.220:1081 | 否 | ✓ 1683ms | ✓ 539ms | 否 | ✓ 1182ms | http |
| 34.71.229.255:3128 | 否 | 否 | ✓ 1777ms | ✓ 1925ms | ✓ 1143ms | http |
| 94.158.219.111:3128 | ✓ 1836ms | 否 | 否 | ✓ 1778ms | ✓ 1541ms | http |
| 168.144.75.9:3128 | ✓ 1314ms | 否 | ✓ 1569ms | ✓ 1899ms | 否 | http |
| 60.249.94.208:3128 | ✓ 909ms | ✓ 1153ms | ✓ 949ms | ✓ 1189ms | ✓ 1546ms | http |
| 129.150.53.35:3128 | ✓ 1398ms | 否 | 否 | ✓ 1431ms | ✓ 1271ms | http |
| 101.32.243.189:80 | ✓ 1451ms | 否 | ✓ 1705ms | ✓ 1799ms | ✓ 1560ms | http |
| 121.43.196.213:8222 | 否 | 否 | ✓ 1125ms | ✓ 1503ms | ✓ 1242ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1728ms | ✓ 1358ms | ✓ 1710ms | 否 | http |
| 91.107.181.137:3128 | ✓ 1029ms | ✓ 1368ms | ✓ 1986ms | ✓ 1978ms | ✓ 1615ms | http |
| 183.232.248.73:7890 | ✓ 1434ms | 否 | ✓ 1083ms | ✓ 1673ms | 否 | http |

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
