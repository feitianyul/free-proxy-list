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

最后更新：2026-03-08 08:41:45 UTC（2026-03-08 16:41:45 UTC+8）

**代理总数：64**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 63 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 64 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 206ms | ✓ 1860ms | ✓ 1101ms | ✓ 1181ms | ✓ 924ms | http |
| 103.84.95.54:7890 | ✓ 762ms | 否 | 否 | ✓ 1082ms | ✓ 1111ms | http |
| 162.240.154.26:3128 | ✓ 846ms | 否 | 否 | ✓ 1112ms | ✓ 871ms | http |
| 1.231.81.166:3128 | ✓ 1956ms | ✓ 1675ms | ✓ 1162ms | ✓ 1192ms | ✓ 1007ms | http |
| 114.4.251.26:8080 | ✓ 1967ms | 否 | ✓ 1344ms | ✓ 1949ms | 否 | http |
| 14.56.107.244:3128 | ✓ 1965ms | 否 | ✓ 1517ms | 否 | ✓ 1778ms | http |
| 62.113.119.14:8080 | ✓ 1409ms | 否 | ✓ 1071ms | ✓ 1673ms | ✓ 1144ms | http |
| 152.42.213.210:8080 | ✓ 802ms | 否 | ✓ 952ms | ✓ 1176ms | ✓ 1077ms | http |
| 162.248.165.72:1080 | ✓ 912ms | 否 | 否 | ✓ 1798ms | ✓ 1269ms | http |
| 35.225.22.61:80 | 否 | 否 | ✓ 496ms | ✓ 949ms | ✓ 999ms | http |
| 117.159.239.49:22222 | ✓ 959ms | ✓ 1098ms | ✓ 939ms | ✓ 1211ms | ✓ 893ms | http |
| 180.191.230.220:8082 | 否 | 否 | ✓ 1584ms | ✓ 1584ms | ✓ 1602ms | http |
| 168.235.110.63:3128 | ✓ 1249ms | ✓ 1309ms | ✓ 1853ms | ✓ 1115ms | ✓ 853ms | http |
| 202.155.12.161:443 | ✓ 1492ms | 否 | ✓ 1279ms | ✓ 1069ms | ✓ 1025ms | http |
| 120.92.212.16:8890 | ✓ 1050ms | ✓ 1554ms | ✓ 1972ms | ✓ 1294ms | 否 | http |
| 193.168.173.136:443 | ✓ 983ms | 否 | ✓ 1008ms | ✓ 1947ms | 否 | http |
| 46.183.25.8:443 | ✓ 1578ms | 否 | ✓ 1294ms | 否 | ✓ 859ms | http |
| 159.89.31.62:8080 | ✓ 818ms | ✓ 1990ms | ✓ 1733ms | ✓ 1959ms | ✓ 1445ms | http |
| 121.128.121.54:3128 | ✓ 1713ms | ✓ 1519ms | 否 | ✓ 1040ms | 否 | http |
| 101.43.255.96:80 | ✓ 1280ms | 否 | ✓ 1058ms | ✓ 1374ms | ✓ 1412ms | http |
| 42.115.72.27:2033 | 否 | 否 | ✓ 1576ms | ✓ 1816ms | ✓ 1788ms | http |
| 120.240.35.175:22222 | ✓ 977ms | ✓ 1353ms | ✓ 1112ms | ✓ 1293ms | ✓ 1014ms | http |
| 61.72.110.54:3128 | ✓ 1695ms | ✓ 1010ms | ✓ 1302ms | ✓ 1140ms | 否 | http |
| 88.80.150.82:8080 | ✓ 1521ms | ✓ 1992ms | 否 | 否 | ✓ 1767ms | https |
| 120.240.35.161:22222 | ✓ 1151ms | ✓ 1273ms | 否 | 否 | ✓ 988ms | http |
| 81.70.169.194:80 | ✓ 1027ms | 否 | 否 | ✓ 1389ms | ✓ 1075ms | http |
| 42.115.72.27:2039 | ✓ 1553ms | 否 | 否 | ✓ 1742ms | ✓ 1773ms | http |
| 222.184.48.241:22222 | ✓ 1046ms | 否 | ✓ 1184ms | ✓ 1343ms | ✓ 1491ms | http |
| 130.110.250.13:1111 | ✓ 1147ms | 否 | ✓ 1983ms | 否 | ✓ 1111ms | http |
| 188.132.141.249:443 | ✓ 1955ms | 否 | ✓ 1966ms | 否 | ✓ 1691ms | http |
| 120.238.159.227:22222 | ✓ 1126ms | ✓ 1281ms | ✓ 1058ms | 否 | 否 | http |
| 115.231.181.40:8128 | 否 | ✓ 1147ms | ✓ 986ms | 否 | ✓ 950ms | http |
| 152.42.213.210:80 | ✓ 991ms | 否 | ✓ 1496ms | ✓ 1203ms | 否 | http |
| 85.9.195.140:1080 | ✓ 1797ms | 否 | ✓ 1103ms | 否 | ✓ 1700ms | http |
| 94.176.3.43:7443 | ✓ 1704ms | 否 | ✓ 1814ms | 否 | ✓ 1837ms | http |
| 183.249.5.213:22222 | ✓ 855ms | ✓ 993ms | 否 | ✓ 1362ms | 否 | http |
| 120.240.35.176:22222 | ✓ 1037ms | ✓ 1359ms | ✓ 1152ms | ✓ 1196ms | ✓ 995ms | http |
| 185.243.218.43:49153 | ✓ 1260ms | 否 | 否 | ✓ 1984ms | ✓ 1638ms | http |
| 39.104.201.40:7890 | 否 | ✓ 1324ms | 否 | ✓ 1317ms | ✓ 1322ms | http |
| 113.59.32.142:22222 | 否 | ✓ 1774ms | ✓ 1453ms | 否 | ✓ 1426ms | http |
| 120.92.212.16:7890 | ✓ 1008ms | 否 | ✓ 1990ms | 否 | ✓ 1254ms | http |
| 173.212.246.157:3128 | ✓ 1524ms | 否 | ✓ 1198ms | ✓ 1981ms | ✓ 1582ms | http |
| 45.136.198.40:3128 | ✓ 1479ms | ✓ 1559ms | ✓ 1493ms | 否 | ✓ 1614ms | http |
| 165.227.5.10:8888 | ✓ 462ms | ✓ 1771ms | ✓ 390ms | ✓ 858ms | ✓ 1731ms | http |
| 65.108.203.36:28080 | ✓ 1397ms | 否 | ✓ 1970ms | 否 | ✓ 1845ms | http |
| 14.225.222.164:7890 | ✓ 1825ms | 否 | 否 | ✓ 1688ms | ✓ 1633ms | http |
| 34.101.184.164:3128 | ✓ 932ms | 否 | ✓ 1316ms | ✓ 1811ms | 否 | http |
| 59.46.216.131:30001 | ✓ 1114ms | ✓ 1413ms | ✓ 1204ms | 否 | 否 | http |
| 103.215.36.88:16345 | ✓ 1105ms | ✓ 1320ms | ✓ 1153ms | 否 | ✓ 1956ms | http |
| 103.183.10.203:3125 | ✓ 1683ms | 否 | 否 | ✓ 1440ms | ✓ 1735ms | http |
| 161.97.131.23:8899 | 否 | 否 | ✓ 990ms | ✓ 1772ms | ✓ 1376ms | http |
| 210.223.44.230:3128 | ✓ 1734ms | ✓ 880ms | ✓ 956ms | ✓ 1014ms | 否 | http |
| 117.159.239.44:22222 | ✓ 867ms | 否 | ✓ 1119ms | ✓ 1221ms | ✓ 938ms | http |
| 157.0.142.246:10057 | ✓ 1106ms | 否 | 否 | ✓ 1406ms | ✓ 1123ms | http |
| 183.249.5.111:22222 | 否 | 否 | ✓ 845ms | ✓ 1093ms | ✓ 772ms | http |
| 120.240.35.177:22222 | 否 | ✓ 1379ms | ✓ 1038ms | 否 | ✓ 1063ms | http |
| 194.213.18.200:443 | ✓ 969ms | 否 | ✓ 1737ms | ✓ 1910ms | ✓ 914ms | http |
| 46.249.103.192:443 | ✓ 640ms | 否 | ✓ 1078ms | ✓ 1690ms | 否 | http |
| 103.67.46.225:3125 | ✓ 1960ms | 否 | ✓ 1880ms | ✓ 1650ms | ✓ 1944ms | http |
| 120.198.141.75:22222 | 否 | ✓ 1626ms | ✓ 1203ms | ✓ 1482ms | ✓ 1208ms | http |
| 113.59.32.141:22222 | ✓ 1312ms | ✓ 1873ms | ✓ 1576ms | ✓ 1819ms | ✓ 1438ms | http |
| 47.77.193.180:1080 | ✓ 934ms | ✓ 1814ms | ✓ 374ms | ✓ 783ms | ✓ 596ms | http |
| 120.240.35.160:22222 | ✓ 1198ms | ✓ 1384ms | ✓ 1057ms | ✓ 1251ms | ✓ 1007ms | http |
| 112.78.187.186:8080 | 否 | 否 | ✓ 1576ms | ✓ 1505ms | ✓ 1474ms | http |

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
