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

最后更新：2026-03-02 16:39:28 UTC（2026-03-03 00:39:28 UTC+8）

**代理总数：65**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 65 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 250ms | ✓ 902ms | ✓ 796ms | 否 | ✓ 1002ms | http |
| 2.56.178.131:443 | ✓ 1025ms | 否 | ✓ 1427ms | 否 | ✓ 1657ms | http |
| 14.56.177.44:3128 | 否 | 否 | ✓ 874ms | ✓ 1376ms | ✓ 1094ms | http |
| 121.128.121.54:3128 | 否 | 否 | ✓ 1338ms | ✓ 1241ms | ✓ 988ms | http |
| 14.56.107.244:3128 | 否 | 否 | ✓ 1221ms | ✓ 1213ms | ✓ 940ms | http |
| 186.148.180.46:999 | ✓ 1110ms | 否 | ✓ 1259ms | ✓ 1836ms | ✓ 1538ms | http |
| 162.240.154.26:3128 | ✓ 792ms | ✓ 1307ms | ✓ 1311ms | ✓ 1424ms | ✓ 1338ms | http |
| 61.72.221.94:3128 | ✓ 1122ms | 否 | 否 | ✓ 1788ms | ✓ 1874ms | http |
| 35.225.22.61:80 | 否 | ✓ 1894ms | ✓ 374ms | 否 | ✓ 1058ms | http |
| 103.84.95.54:7890 | ✓ 840ms | 否 | 否 | ✓ 1744ms | ✓ 876ms | http |
| 210.223.44.230:3128 | ✓ 778ms | ✓ 1580ms | ✓ 1342ms | ✓ 1141ms | ✓ 920ms | http |
| 45.125.67.37:8443 | ✓ 1151ms | 否 | ✓ 1117ms | 否 | ✓ 1424ms | http |
| 61.72.110.54:3128 | ✓ 1649ms | ✓ 1838ms | 否 | 否 | ✓ 1272ms | http |
| 172.212.68.37:3128 | ✓ 304ms | 否 | ✓ 1542ms | ✓ 1002ms | ✓ 1140ms | http |
| 165.227.5.10:8888 | ✓ 920ms | 否 | ✓ 1855ms | ✓ 1395ms | ✓ 1024ms | http |
| 222.228.171.92:8080 | ✓ 1421ms | 否 | ✓ 1455ms | ✓ 1365ms | ✓ 1038ms | http |
| 91.238.104.172:2024 | ✓ 1508ms | ✓ 1478ms | ✓ 1936ms | 否 | ✓ 1756ms | http |
| 195.123.209.48:3128 | ✓ 1514ms | ✓ 1581ms | 否 | ✓ 1902ms | ✓ 1756ms | http |
| 91.238.104.171:2023 | ✓ 1512ms | 否 | ✓ 1412ms | 否 | ✓ 1894ms | http |
| 120.92.212.16:8890 | ✓ 1162ms | ✓ 1460ms | 否 | 否 | ✓ 1147ms | http |
| 101.43.255.96:80 | ✓ 1618ms | 否 | ✓ 1386ms | 否 | ✓ 1284ms | http |
| 35.234.17.221:8080 | 否 | ✓ 1467ms | ✓ 1492ms | 否 | ✓ 1048ms | http |
| 120.92.212.16:7890 | ✓ 1784ms | ✓ 1409ms | ✓ 1805ms | 否 | 否 | http |
| 91.233.223.147:3128 | ✓ 764ms | ✓ 1876ms | ✓ 1242ms | ✓ 1833ms | ✓ 1461ms | http |
| 34.101.184.164:3128 | ✓ 1757ms | 否 | ✓ 1099ms | ✓ 1473ms | ✓ 1163ms | http |
| 115.231.181.40:8128 | ✓ 1036ms | 否 | 否 | ✓ 1813ms | ✓ 1098ms | http |
| 103.139.138.194:3128 | ✓ 1771ms | 否 | ✓ 1628ms | ✓ 1626ms | ✓ 1450ms | http |
| 61.72.110.94:3128 | ✓ 795ms | 否 | ✓ 1715ms | 否 | ✓ 1831ms | http |
| 61.109.216.213:8080 | ✓ 1153ms | 否 | ✓ 1329ms | ✓ 1487ms | ✓ 1261ms | http |
| 138.124.53.25:7443 | ✓ 415ms | ✓ 1998ms | ✓ 1730ms | ✓ 1511ms | 否 | http |
| 125.128.12.194:3128 | ✓ 1726ms | ✓ 1492ms | ✓ 1739ms | 否 | 否 | http |
| 125.128.12.14:3128 | 否 | 否 | ✓ 1764ms | ✓ 1906ms | ✓ 1330ms | http |
| 125.128.12.114:3128 | ✓ 1702ms | 否 | ✓ 1650ms | ✓ 1629ms | 否 | http |
| 115.76.5.32:10010 | 否 | 否 | ✓ 1703ms | ✓ 1804ms | ✓ 1821ms | http |
| 5.75.196.26:40000 | ✓ 437ms | ✓ 1399ms | ✓ 1580ms | 否 | 否 | http |
| 187.216.141.46:3128 | ✓ 574ms | 否 | ✓ 1002ms | 否 | ✓ 1068ms | http |
| 209.38.51.97:3128 | 否 | ✓ 1865ms | ✓ 964ms | 否 | ✓ 931ms | http |
| 125.128.12.144:3128 | ✓ 1375ms | 否 | ✓ 1907ms | 否 | ✓ 1777ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1604ms | 否 | ✓ 1838ms | ✓ 1866ms | http |
| 95.85.252.153:21064 | ✓ 612ms | ✓ 1479ms | ✓ 1599ms | 否 | 否 | http |
| 175.215.38.158:3024 | ✓ 1403ms | 否 | ✓ 1988ms | 否 | ✓ 1876ms | http |
| 110.172.29.131:3128 | 否 | 否 | ✓ 1920ms | ✓ 1407ms | ✓ 1432ms | http |
| 20.2.83.243:3128 | ✓ 860ms | ✓ 1415ms | ✓ 938ms | ✓ 1245ms | ✓ 1276ms | http |
| 81.70.169.194:80 | ✓ 1498ms | ✓ 1968ms | ✓ 1211ms | 否 | 否 | http |
| 90.84.188.97:8000 | ✓ 519ms | ✓ 1616ms | 否 | 否 | ✓ 1688ms | http |
| 45.129.141.143:3128 | 否 | 否 | ✓ 1779ms | ✓ 1909ms | ✓ 1519ms | http |
| 45.136.198.40:3128 | ✓ 1240ms | 否 | ✓ 1709ms | 否 | ✓ 1716ms | http |
| 45.140.147.82:1081 | 否 | ✓ 1510ms | ✓ 941ms | ✓ 1743ms | ✓ 1271ms | http |
| 183.128.208.68:7890 | ✓ 1076ms | ✓ 1349ms | ✓ 1165ms | 否 | 否 | http |
| 115.76.5.32:10006 | ✓ 1600ms | 否 | 否 | ✓ 1859ms | ✓ 1540ms | http |
| 5.129.228.225:1080 | ✓ 1371ms | 否 | ✓ 1256ms | ✓ 1262ms | 否 | http |
| 115.76.5.32:10009 | 否 | 否 | ✓ 1686ms | ✓ 1894ms | ✓ 1565ms | http |
| 223.16.170.103:3128 | ✓ 1127ms | 否 | 否 | ✓ 1345ms | ✓ 1372ms | http |
| 142.171.85.32:1080 | ✓ 1032ms | 否 | 否 | ✓ 1786ms | ✓ 1386ms | http |
| 120.79.99.232:8099 | ✓ 1434ms | ✓ 1760ms | ✓ 1526ms | ✓ 1680ms | ✓ 1429ms | http |
| 47.110.42.192:9003 | ✓ 1726ms | ✓ 1638ms | ✓ 1579ms | ✓ 1887ms | ✓ 1926ms | http |
| 113.165.202.31:1002 | ✓ 1694ms | 否 | ✓ 1633ms | ✓ 1839ms | ✓ 1678ms | http |
| 103.179.218.14:8000 | ✓ 1939ms | 否 | ✓ 1876ms | ✓ 1654ms | ✓ 1652ms | http |
| 150.249.255.91:3128 | 否 | 否 | ✓ 784ms | ✓ 1066ms | ✓ 863ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1501ms | ✓ 1558ms | ✓ 1691ms | http |
| 45.140.147.82:1082 | ✓ 518ms | 否 | ✓ 1032ms | ✓ 1543ms | ✓ 1991ms | http |
| 46.249.103.192:443 | ✓ 954ms | 否 | ✓ 1312ms | ✓ 1975ms | 否 | http |
| 37.27.100.102:443 | ✓ 603ms | 否 | ✓ 1849ms | 否 | ✓ 1860ms | http |
| 103.74.192.243:7890 | 否 | ✓ 1946ms | ✓ 1076ms | ✓ 1728ms | ✓ 1711ms | http |
| 47.77.180.205:1080 | 否 | 否 | ✓ 593ms | ✓ 968ms | ✓ 647ms | http |

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
