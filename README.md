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

最后更新：2026-02-28 15:30:56 UTC（2026-02-28 23:30:56 UTC+8）

**代理总数：82**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 82 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 3.213.157.4:3128 | ✓ 189ms | ✓ 1332ms | ✓ 1062ms | ✓ 1408ms | ✓ 1096ms | http |
| 205.209.118.30:3138 | ✓ 341ms | ✓ 1952ms | ✓ 1806ms | ✓ 1081ms | ✓ 830ms | http |
| 45.140.147.155:1081 | ✓ 1498ms | ✓ 1958ms | ✓ 1403ms | 否 | 否 | http |
| 193.124.225.175:1080 | ✓ 1152ms | 否 | ✓ 1584ms | ✓ 1939ms | 否 | http |
| 120.92.212.16:8890 | ✓ 1445ms | ✓ 1504ms | 否 | ✓ 1767ms | 否 | http |
| 35.225.22.61:80 | ✓ 794ms | ✓ 1382ms | 否 | ✓ 1015ms | 否 | http |
| 101.43.255.96:80 | ✓ 1121ms | 否 | 否 | ✓ 1581ms | ✓ 1248ms | http |
| 81.70.169.194:80 | ✓ 1135ms | 否 | ✓ 1579ms | ✓ 1512ms | 否 | http |
| 185.115.74.185:8080 | ✓ 982ms | ✓ 1748ms | ✓ 1954ms | 否 | 否 | http |
| 138.124.53.25:7443 | ✓ 387ms | 否 | ✓ 1290ms | 否 | ✓ 1933ms | http |
| 165.225.120.17:10458 | ✓ 1896ms | 否 | ✓ 1614ms | 否 | ✓ 1512ms | http |
| 165.225.120.17:10880 | ✓ 1888ms | 否 | ✓ 1795ms | 否 | ✓ 1349ms | http |
| 165.225.120.17:11745 | ✓ 1889ms | 否 | ✓ 1819ms | 否 | ✓ 1339ms | http |
| 165.225.120.17:11912 | ✓ 1894ms | 否 | ✓ 1575ms | 否 | ✓ 1580ms | http |
| 165.225.120.17:11995 | ✓ 1892ms | 否 | ✓ 1614ms | 否 | ✓ 1621ms | http |
| 165.225.120.17:10906 | ✓ 1899ms | 否 | ✓ 1825ms | 否 | ✓ 1340ms | http |
| 165.225.120.17:11702 | ✓ 1897ms | 否 | ✓ 1623ms | 否 | ✓ 1620ms | http |
| 136.226.254.24:11933 | ✓ 1890ms | 否 | ✓ 1643ms | 否 | ✓ 1684ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1731ms | ✓ 1441ms | ✓ 1460ms | 否 | http |
| 165.225.120.17:11178 | ✓ 1891ms | 否 | ✓ 1824ms | 否 | ✓ 1481ms | http |
| 136.226.254.24:11819 | ✓ 1893ms | 否 | ✓ 1805ms | ✓ 1982ms | 否 | http |
| 37.187.109.70:10111 | ✓ 1262ms | ✓ 1225ms | ✓ 1765ms | 否 | 否 | http |
| 162.240.154.26:3128 | ✓ 1777ms | 否 | ✓ 1172ms | ✓ 1823ms | 否 | http |
| 45.140.147.82:1081 | ✓ 1549ms | 否 | ✓ 678ms | ✓ 1266ms | ✓ 975ms | http |
| 142.171.85.32:1080 | ✓ 1069ms | ✓ 1297ms | 否 | ✓ 1319ms | ✓ 1594ms | http |
| 165.225.120.17:12497 | ✓ 831ms | 否 | ✓ 855ms | ✓ 1762ms | ✓ 1372ms | http |
| 165.225.120.17:12265 | ✓ 834ms | 否 | ✓ 832ms | ✓ 1719ms | ✓ 1328ms | http |
| 165.225.120.17:11070 | ✓ 871ms | 否 | ✓ 830ms | ✓ 1740ms | 否 | http |
| 45.140.147.82:1082 | ✓ 443ms | ✓ 1640ms | ✓ 1559ms | ✓ 1638ms | 否 | http |
| 165.225.120.17:11099 | ✓ 846ms | 否 | ✓ 802ms | ✓ 1775ms | ✓ 1356ms | http |
| 165.225.120.17:12215 | ✓ 836ms | 否 | ✓ 827ms | ✓ 1768ms | ✓ 1382ms | http |
| 165.225.120.17:10919 | ✓ 867ms | 否 | ✓ 810ms | ✓ 1782ms | 否 | http |
| 62.113.119.14:8080 | ✓ 704ms | 否 | ✓ 1106ms | ✓ 1468ms | ✓ 1172ms | http |
| 168.235.110.63:3128 | ✓ 1257ms | 否 | ✓ 734ms | ✓ 1199ms | ✓ 748ms | http |
| 115.231.181.40:8128 | ✓ 1936ms | ✓ 1408ms | ✓ 1392ms | 否 | 否 | http |
| 103.113.70.189:1081 | 否 | ✓ 1153ms | 否 | ✓ 1073ms | ✓ 788ms | http |
| 136.226.254.24:11197 | ✓ 1114ms | 否 | ✓ 1798ms | 否 | ✓ 1497ms | http |
| 121.237.181.137:8888 | ✓ 1930ms | 否 | ✓ 1047ms | ✓ 1804ms | ✓ 1108ms | http |
| 104.238.30.68:63744 | ✓ 1715ms | 否 | ✓ 1807ms | 否 | ✓ 1935ms | http |
| 104.238.30.63:63744 | ✓ 1619ms | 否 | ✓ 1903ms | 否 | ✓ 1839ms | http |
| 121.230.8.34:1080 | ✓ 1447ms | 否 | ✓ 1220ms | ✓ 1750ms | ✓ 1353ms | http |
| 121.230.8.220:1080 | ✓ 1395ms | 否 | 否 | ✓ 1867ms | ✓ 1312ms | http |
| 101.32.244.83:8080 | ✓ 1199ms | 否 | ✓ 1183ms | ✓ 1758ms | ✓ 1533ms | http |
| 121.43.196.213:8222 | ✓ 1143ms | ✓ 1245ms | ✓ 1064ms | ✓ 1364ms | ✓ 1089ms | http |
| 121.43.196.210:8222 | ✓ 1145ms | ✓ 1272ms | ✓ 1030ms | 否 | ✓ 1058ms | http |
| 114.55.226.123:10086 | ✓ 1237ms | 否 | ✓ 1455ms | ✓ 1559ms | ✓ 1267ms | http |
| 18.233.58.30:80 | ✓ 154ms | 否 | ✓ 773ms | 否 | ✓ 667ms | http |
| 45.125.67.37:8443 | 否 | 否 | ✓ 1161ms | ✓ 1294ms | ✓ 1131ms | http |
| 188.166.208.168:9876 | ✓ 1742ms | 否 | ✓ 957ms | ✓ 1308ms | ✓ 1057ms | http |
| 136.226.254.24:12265 | ✓ 1771ms | 否 | ✓ 984ms | ✓ 1817ms | ✓ 1355ms | http |
| 136.226.254.24:11909 | ✓ 1771ms | 否 | ✓ 992ms | ✓ 1913ms | ✓ 1393ms | http |
| 136.226.254.24:12360 | ✓ 1769ms | 否 | ✓ 990ms | ✓ 1888ms | ✓ 1396ms | http |
| 104.238.30.45:59741 | ✓ 1658ms | 否 | ✓ 1711ms | 否 | ✓ 1931ms | http |
| 136.226.254.24:11342 | ✓ 1774ms | 否 | ✓ 989ms | ✓ 1863ms | 否 | http |
| 165.225.120.17:11943 | ✓ 1774ms | 否 | ✓ 1722ms | 否 | ✓ 1362ms | http |
| 136.226.254.24:10517 | ✓ 1779ms | 否 | ✓ 988ms | 否 | ✓ 1395ms | http |
| 136.226.254.24:10919 | ✓ 1778ms | 否 | ✓ 986ms | ✓ 1807ms | 否 | http |
| 34.159.121.205:3128 | ✓ 896ms | 否 | ✓ 925ms | ✓ 1671ms | ✓ 1351ms | http |
| 34.7.88.87:3128 | ✓ 904ms | 否 | ✓ 1360ms | ✓ 1859ms | ✓ 1357ms | http |
| 34.79.102.160:3128 | ✓ 906ms | ✓ 1635ms | ✓ 1199ms | ✓ 1574ms | ✓ 1289ms | http |
| 104.238.30.40:59741 | ✓ 1649ms | 否 | ✓ 1683ms | 否 | ✓ 1931ms | http |
| 34.32.154.33:3128 | ✓ 466ms | ✓ 1687ms | ✓ 780ms | ✓ 1656ms | ✓ 1379ms | http |
| 34.78.177.18:3128 | ✓ 450ms | 否 | ✓ 1034ms | ✓ 1618ms | ✓ 1445ms | http |
| 34.78.200.22:3128 | ✓ 449ms | 否 | ✓ 907ms | 否 | ✓ 1299ms | http |
| 35.204.245.176:3128 | ✓ 434ms | 否 | ✓ 968ms | ✓ 1876ms | ✓ 1362ms | http |
| 35.241.222.101:3128 | ✓ 447ms | 否 | ✓ 1239ms | ✓ 1563ms | ✓ 1738ms | http |
| 103.84.95.54:7890 | ✓ 1431ms | 否 | ✓ 1456ms | 否 | ✓ 1192ms | http |
| 36.147.78.166:80 | 否 | ✓ 1906ms | 否 | ✓ 1859ms | ✓ 1933ms | http |
| 36.147.78.166:443 | ✓ 1920ms | ✓ 1897ms | ✓ 1983ms | ✓ 1890ms | ✓ 1912ms | http |
| 104.238.30.37:59741 | ✓ 1587ms | 否 | ✓ 1903ms | 否 | ✓ 1875ms | http |
| 34.185.159.217:3128 | ✓ 966ms | 否 | ✓ 1162ms | ✓ 1777ms | ✓ 1470ms | http |
| 104.238.30.39:59741 | ✓ 1614ms | 否 | ✓ 1872ms | 否 | ✓ 1839ms | http |
| 104.238.30.38:59741 | ✓ 1561ms | 否 | ✓ 1903ms | 否 | ✓ 1871ms | http |
| 104.238.30.50:59741 | ✓ 1820ms | 否 | ✓ 1679ms | 否 | ✓ 1903ms | http |
| 45.136.198.40:3128 | ✓ 659ms | ✓ 1637ms | ✓ 1469ms | 否 | ✓ 1596ms | http |
| 104.238.30.86:63900 | ✓ 1731ms | 否 | ✓ 1712ms | 否 | ✓ 1934ms | http |
| 45.177.178.242:999 | 否 | ✓ 1865ms | ✓ 1226ms | ✓ 1179ms | ✓ 1534ms | http |
| 45.129.141.143:3128 | ✓ 1632ms | 否 | ✓ 1453ms | ✓ 1900ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1527ms | ✓ 1636ms | ✓ 1857ms | 否 | 否 | http |
| 35.234.17.221:8080 | 否 | ✓ 1744ms | ✓ 1309ms | ✓ 1236ms | 否 | http |
| 217.76.245.80:999 | ✓ 598ms | 否 | ✓ 1050ms | ✓ 1334ms | ✓ 1391ms | http |
| 180.127.149.245:1080 | ✓ 1139ms | 否 | ✓ 1157ms | 否 | ✓ 1532ms | http |

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
