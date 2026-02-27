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

最后更新：2026-02-27 06:52:27 UTC（2026-02-27 14:52:27 UTC+8）

**代理总数：61**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 61 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 103.84.95.54:7890 | ✓ 839ms | 否 | 否 | ✓ 973ms | ✓ 1016ms | http |
| 101.47.73.135:3128 | ✓ 1216ms | 否 | ✓ 1377ms | ✓ 1972ms | ✓ 1256ms | http |
| 138.124.53.25:7443 | ✓ 502ms | 否 | ✓ 1653ms | ✓ 1639ms | ✓ 1355ms | http |
| 120.92.212.16:8890 | ✓ 1016ms | 否 | ✓ 1080ms | ✓ 1390ms | 否 | http |
| 211.230.49.122:3128 | ✓ 891ms | ✓ 1583ms | 否 | ✓ 1570ms | ✓ 1376ms | http |
| 147.45.216.148:1080 | ✓ 637ms | 否 | ✓ 1510ms | ✓ 1920ms | ✓ 1558ms | http |
| 185.246.90.163:10808 | ✓ 713ms | ✓ 1836ms | 否 | 否 | ✓ 1684ms | http |
| 72.56.59.62:63133 | ✓ 1723ms | 否 | ✓ 1992ms | 否 | ✓ 1968ms | http |
| 72.56.59.23:61937 | ✓ 1760ms | 否 | ✓ 1737ms | 否 | ✓ 1966ms | http |
| 81.177.48.54:2080 | ✓ 1125ms | 否 | ✓ 1882ms | 否 | ✓ 1751ms | http |
| 35.225.22.61:80 | ✓ 1153ms | 否 | ✓ 1013ms | ✓ 959ms | ✓ 812ms | http |
| 120.92.212.16:7890 | ✓ 1046ms | 否 | 否 | ✓ 1548ms | ✓ 1061ms | http |
| 72.56.50.17:59787 | ✓ 1697ms | 否 | ✓ 1837ms | 否 | ✓ 1964ms | http |
| 103.236.64.247:8888 | ✓ 1182ms | ✓ 1863ms | 否 | ✓ 1449ms | 否 | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1706ms | ✓ 1418ms | ✓ 1191ms | http |
| 124.16.93.70:7890 | ✓ 1008ms | 否 | ✓ 1118ms | ✓ 1244ms | 否 | http |
| 210.223.44.230:3128 | ✓ 1948ms | 否 | 否 | ✓ 1195ms | ✓ 979ms | http |
| 14.56.107.244:3128 | ✓ 1806ms | 否 | ✓ 1573ms | ✓ 1112ms | 否 | http |
| 72.56.59.17:61931 | ✓ 1745ms | 否 | ✓ 1842ms | 否 | ✓ 1997ms | http |
| 170.246.183.250:3128 | ✓ 793ms | 否 | ✓ 796ms | 否 | ✓ 1868ms | http |
| 81.70.169.194:80 | ✓ 1297ms | ✓ 1292ms | ✓ 1056ms | ✓ 1646ms | ✓ 1220ms | http |
| 120.46.152.136:3128 | 否 | ✓ 1978ms | ✓ 1702ms | ✓ 1931ms | ✓ 1748ms | http |
| 35.234.17.221:8080 | ✓ 988ms | 否 | ✓ 1452ms | 否 | ✓ 1178ms | http |
| 168.235.110.63:3128 | ✓ 297ms | ✓ 1050ms | ✓ 843ms | ✓ 1161ms | ✓ 870ms | http |
| 36.147.78.166:80 | ✓ 1877ms | 否 | ✓ 1749ms | ✓ 1928ms | 否 | http |
| 195.123.209.48:3128 | ✓ 762ms | ✓ 1865ms | ✓ 1810ms | 否 | ✓ 1808ms | http |
| 103.139.138.194:3128 | ✓ 1876ms | 否 | ✓ 1559ms | ✓ 1571ms | ✓ 1396ms | http |
| 101.43.255.96:80 | ✓ 999ms | ✓ 1394ms | 否 | ✓ 1639ms | 否 | http |
| 45.125.67.37:8443 | ✓ 1078ms | 否 | 否 | ✓ 1256ms | ✓ 1065ms | http |
| 194.87.43.49:8888 | ✓ 1113ms | ✓ 1984ms | ✓ 1586ms | 否 | 否 | http |
| 91.238.104.155:2023 | ✓ 1392ms | ✓ 1558ms | ✓ 716ms | 否 | ✓ 1248ms | http |
| 72.56.59.56:63127 | ✓ 1674ms | 否 | ✓ 1743ms | 否 | ✓ 1966ms | http |
| 61.72.110.54:3128 | ✓ 1487ms | ✓ 1503ms | ✓ 1175ms | 否 | 否 | http |
| 34.101.184.164:3128 | ✓ 1728ms | 否 | ✓ 997ms | ✓ 1513ms | ✓ 1387ms | http |
| 45.136.198.40:3128 | ✓ 1037ms | 否 | 否 | ✓ 1994ms | ✓ 1647ms | http |
| 103.104.99.29:80 | ✓ 1695ms | 否 | ✓ 1596ms | ✓ 1607ms | ✓ 1427ms | http |
| 103.104.99.89:80 | ✓ 1702ms | 否 | ✓ 1652ms | ✓ 1452ms | ✓ 1712ms | http |
| 61.72.110.94:3128 | ✓ 909ms | 否 | ✓ 1132ms | ✓ 1334ms | 否 | http |
| 121.237.181.137:8888 | ✓ 1035ms | ✓ 1158ms | ✓ 930ms | ✓ 1342ms | ✓ 897ms | http |
| 111.79.111.126:3128 | 否 | 否 | ✓ 1737ms | ✓ 1547ms | ✓ 1402ms | http |
| 14.56.118.34:3128 | ✓ 1816ms | 否 | 否 | ✓ 1739ms | ✓ 1519ms | http |
| 61.72.110.24:3128 | ✓ 1588ms | ✓ 1635ms | 否 | ✓ 1718ms | 否 | http |
| 202.129.206.239:3128 | 否 | 否 | ✓ 1644ms | ✓ 1606ms | ✓ 1366ms | http |
| 91.238.105.64:2024 | 否 | 否 | ✓ 1338ms | ✓ 1863ms | ✓ 1958ms | http |
| 103.149.99.128:3128 | 否 | 否 | ✓ 1527ms | ✓ 1191ms | ✓ 1427ms | http |
| 115.231.181.40:8128 | 否 | ✓ 1140ms | ✓ 1037ms | ✓ 1977ms | ✓ 1059ms | http |
| 152.32.255.24:27197 | ✓ 1873ms | 否 | 否 | ✓ 1520ms | ✓ 1055ms | http |
| 52.188.28.218:3128 | ✓ 1147ms | ✓ 1920ms | ✓ 1773ms | 否 | 否 | http |
| 45.140.147.155:1082 | ✓ 504ms | ✓ 1495ms | ✓ 1690ms | 否 | ✓ 1050ms | http |
| 45.140.147.155:1081 | ✓ 551ms | 否 | ✓ 1143ms | ✓ 1733ms | ✓ 1280ms | http |
| 8.219.97.248:80 | 否 | 否 | ✓ 1842ms | ✓ 1632ms | ✓ 1888ms | http |
| 37.27.100.107:443 | ✓ 1117ms | 否 | ✓ 1645ms | ✓ 1911ms | 否 | http |
| 172.212.68.37:3128 | ✓ 826ms | ✓ 1673ms | ✓ 639ms | ✓ 1558ms | ✓ 1091ms | http |
| 36.147.78.166:443 | ✓ 1818ms | 否 | ✓ 1791ms | 否 | ✓ 1712ms | http |
| 103.39.51.190:8080 | ✓ 1775ms | 否 | 否 | ✓ 1695ms | ✓ 1648ms | http |
| 18.100.127.30:3128 | ✓ 1640ms | 否 | ✓ 1484ms | 否 | ✓ 1630ms | http |
| 62.113.119.14:8080 | ✓ 1430ms | 否 | ✓ 1547ms | 否 | ✓ 1154ms | http |
| 223.16.170.103:3128 | 否 | 否 | ✓ 1188ms | ✓ 1199ms | ✓ 1207ms | http |
| 85.208.108.43:2094 | ✓ 616ms | 否 | ✓ 969ms | ✓ 1070ms | ✓ 752ms | http |
| 85.208.108.43:10808 | ✓ 792ms | 否 | ✓ 153ms | ✓ 1048ms | ✓ 737ms | http |
| 45.10.69.30:8888 | 否 | ✓ 1412ms | ✓ 448ms | ✓ 953ms | 否 | http |

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
