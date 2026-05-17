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

最后更新：2026-05-17 14:30:38 UTC（2026-05-17 22:30:38 UTC+8）

**代理总数：60**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 60 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 113.160.132.26:8080 | ✓ 1807ms | ✓ 1564ms | ✓ 1124ms | ✓ 1313ms | ✓ 1091ms | http |
| 65.109.125.111:8443 | ✓ 1237ms | 否 | ✓ 1784ms | 否 | ✓ 1906ms | http |
| 185.200.188.234:10001 | ✓ 1520ms | 否 | ✓ 1771ms | ✓ 1987ms | ✓ 1715ms | http |
| 170.106.136.181:31002 | ✓ 954ms | ✓ 677ms | ✓ 589ms | ✓ 749ms | ✓ 579ms | http |
| 115.231.181.40:8128 | ✓ 1322ms | ✓ 1444ms | ✓ 1294ms | ✓ 1314ms | 否 | http |
| 43.156.90.221:10808 | 否 | 否 | ✓ 805ms | ✓ 1455ms | ✓ 1605ms | http |
| 5.252.33.13:2025 | ✓ 1419ms | 否 | ✓ 1328ms | 否 | ✓ 1848ms | http |
| 129.80.238.83:444 | ✓ 432ms | 否 | ✓ 202ms | ✓ 1014ms | ✓ 928ms | http |
| 129.80.217.21:444 | ✓ 448ms | ✓ 1150ms | ✓ 951ms | ✓ 1031ms | ✓ 988ms | http |
| 51.161.50.166:3128 | ✓ 383ms | ✓ 1949ms | ✓ 916ms | ✓ 1230ms | ✓ 1087ms | http |
| 59.46.216.131:30001 | ✓ 1165ms | 否 | ✓ 1333ms | ✓ 1438ms | ✓ 1227ms | http |
| 218.108.131.186:17890 | ✓ 904ms | ✓ 1126ms | ✓ 893ms | ✓ 1205ms | ✓ 959ms | http |
| 34.96.238.40:8080 | 否 | ✓ 1501ms | ✓ 1194ms | 否 | ✓ 1604ms | http |
| 103.21.220.141:3128 | ✓ 844ms | 否 | ✓ 825ms | ✓ 1038ms | ✓ 887ms | http |
| 45.125.67.37:8443 | 否 | 否 | ✓ 1081ms | ✓ 1514ms | ✓ 1832ms | http |
| 128.199.113.85:9090 | ✓ 1868ms | 否 | ✓ 1559ms | ✓ 1164ms | ✓ 1194ms | http |
| 114.214.165.78:10810 | ✓ 1514ms | ✓ 1822ms | ✓ 1666ms | ✓ 1488ms | ✓ 1954ms | http |
| 148.230.4.241:999 | ✓ 863ms | 否 | ✓ 1616ms | ✓ 1417ms | ✓ 1236ms | http |
| 20.164.75.153:8080 | ✓ 1310ms | 否 | ✓ 1794ms | 否 | ✓ 1715ms | http |
| 128.199.116.219:9090 | ✓ 1948ms | 否 | ✓ 1232ms | ✓ 1440ms | ✓ 1002ms | http |
| 185.100.159.138:1080 | ✓ 979ms | 否 | ✓ 1619ms | ✓ 1498ms | ✓ 1257ms | http |
| 8.154.21.175:3128 | ✓ 1406ms | 否 | 否 | ✓ 1155ms | ✓ 967ms | http |
| 34.101.184.164:3128 | ✓ 1818ms | 否 | 否 | ✓ 1757ms | ✓ 1068ms | http |
| 64.188.77.26:3128 | ✓ 1009ms | 否 | ✓ 909ms | ✓ 1784ms | ✓ 1621ms | http |
| 91.242.229.129:8092 | ✓ 1933ms | ✓ 1969ms | ✓ 1594ms | 否 | 否 | http |
| 64.188.77.221:3128 | ✓ 1222ms | 否 | 否 | ✓ 1517ms | ✓ 1344ms | http |
| 138.2.239.213:10010 | ✓ 1017ms | 否 | ✓ 1769ms | ✓ 1209ms | ✓ 1109ms | http |
| 84.47.150.125:1080 | ✓ 1136ms | 否 | ✓ 1576ms | ✓ 1712ms | 否 | http |
| 128.199.121.61:9090 | ✓ 886ms | 否 | 否 | ✓ 1420ms | ✓ 1010ms | http |
| 20.78.118.91:8561 | ✓ 1412ms | ✓ 1307ms | ✓ 577ms | ✓ 925ms | ✓ 736ms | http |
| 20.78.26.206:8561 | ✓ 1413ms | ✓ 967ms | ✓ 829ms | ✓ 1012ms | ✓ 742ms | http |
| 20.210.39.153:8561 | ✓ 1411ms | ✓ 1884ms | ✓ 556ms | ✓ 894ms | ✓ 702ms | http |
| 129.154.225.163:8100 | ✓ 1675ms | 否 | 否 | ✓ 1788ms | ✓ 1743ms | http |
| 152.70.91.193:40000 | ✓ 1851ms | 否 | 否 | ✓ 1996ms | ✓ 1844ms | http |
| 42.96.16.158:1311 | ✓ 1749ms | 否 | ✓ 1157ms | ✓ 1335ms | ✓ 1444ms | http |
| 212.58.132.5:8888 | ✓ 1543ms | 否 | ✓ 1620ms | ✓ 1585ms | ✓ 1295ms | http |
| 193.160.209.58:1080 | ✓ 1400ms | 否 | ✓ 1912ms | 否 | ✓ 1972ms | http |
| 166.88.55.83:7890 | ✓ 735ms | ✓ 1260ms | ✓ 723ms | ✓ 907ms | ✓ 719ms | http |
| 121.130.177.28:8888 | ✓ 1808ms | 否 | ✓ 1840ms | ✓ 1649ms | 否 | http |
| 104.248.151.93:9090 | ✓ 804ms | 否 | ✓ 898ms | ✓ 1164ms | ✓ 921ms | http |
| 159.223.41.216:9090 | ✓ 796ms | 否 | ✓ 906ms | ✓ 1201ms | ✓ 910ms | http |
| 146.190.80.158:9090 | ✓ 1058ms | 否 | ✓ 1032ms | ✓ 1506ms | ✓ 1067ms | http |
| 128.199.114.189:9090 | ✓ 839ms | 否 | ✓ 998ms | ✓ 1251ms | 否 | http |
| 180.125.216.109:8118 | 否 | ✓ 1203ms | ✓ 1123ms | ✓ 1349ms | ✓ 1038ms | http |
| 3.101.133.120:80 | 否 | ✓ 1311ms | ✓ 1395ms | ✓ 1040ms | ✓ 1136ms | http |
| 103.157.78.85:8080 | ✓ 1371ms | 否 | ✓ 1307ms | ✓ 1559ms | ✓ 1501ms | http |
| 147.45.78.89:1080 | ✓ 1654ms | 否 | ✓ 1713ms | 否 | ✓ 1431ms | http |
| 114.214.170.41:27890 | ✓ 1178ms | ✓ 1494ms | ✓ 1451ms | ✓ 1440ms | ✓ 1198ms | http |
| 190.12.150.244:999 | ✓ 1260ms | ✓ 1822ms | ✓ 973ms | 否 | 否 | http |
| 152.32.132.190:7890 | ✓ 921ms | ✓ 1357ms | 否 | ✓ 1470ms | 否 | http |
| 120.92.212.16:7890 | 否 | ✓ 1318ms | ✓ 1271ms | ✓ 1368ms | 否 | http |
| 57.129.144.178:40000 | ✓ 812ms | 否 | ✓ 1169ms | 否 | ✓ 1875ms | http |
| 5.129.248.58:3128 | ✓ 1132ms | 否 | ✓ 1856ms | 否 | ✓ 1711ms | http |
| 210.223.44.230:3128 | 否 | ✓ 1196ms | ✓ 1507ms | ✓ 1462ms | ✓ 1996ms | http |
| 38.211.245.83:999 | ✓ 1445ms | 否 | ✓ 1120ms | 否 | ✓ 1981ms | http |
| 154.12.231.32:80 | ✓ 711ms | 否 | 否 | ✓ 1450ms | ✓ 964ms | http |
| 61.52.131.172:8443 | ✓ 939ms | ✓ 1246ms | ✓ 1140ms | ✓ 1320ms | ✓ 1071ms | http |
| 103.172.70.173:8080 | ✓ 1839ms | 否 | ✓ 1509ms | ✓ 1527ms | ✓ 1461ms | http |
| 152.42.170.187:9090 | ✓ 1232ms | 否 | ✓ 1377ms | ✓ 1446ms | 否 | http |
| 103.147.152.12:1095 | ✓ 1331ms | 否 | ✓ 1352ms | 否 | ✓ 1643ms | http |

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
