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

最后更新：2026-04-24 04:10:49 UTC（2026-04-24 12:10:49 UTC+8）

**代理总数：58**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 58 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 58 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 46.101.95.183:8888 | ✓ 1170ms | 否 | ✓ 784ms | ✓ 1472ms | ✓ 1148ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1322ms | ✓ 969ms | ✓ 1177ms | ✓ 916ms | http |
| 212.58.132.5:8888 | ✓ 1588ms | 否 | ✓ 1556ms | ✓ 1466ms | 否 | http |
| 85.190.99.143:443 | ✓ 1100ms | ✓ 1719ms | ✓ 1511ms | 否 | 否 | http |
| 113.160.132.26:8080 | ✓ 1798ms | ✓ 1918ms | ✓ 1696ms | ✓ 1455ms | ✓ 1612ms | http |
| 47.85.51.197:1080 | ✓ 338ms | ✓ 1122ms | ✓ 278ms | ✓ 990ms | ✓ 757ms | http |
| 177.93.132.244:3128 | ✓ 1051ms | 否 | ✓ 663ms | 否 | ✓ 1704ms | http |
| 159.223.225.118:8888 | 否 | ✓ 1875ms | ✓ 1241ms | ✓ 1549ms | ✓ 1252ms | http |
| 84.47.150.125:1080 | ✓ 1541ms | 否 | ✓ 478ms | ✓ 1954ms | 否 | http |
| 35.225.22.61:80 | ✓ 308ms | ✓ 1495ms | 否 | 否 | ✓ 1320ms | http |
| 91.99.15.45:2095 | ✓ 853ms | ✓ 1586ms | ✓ 1676ms | 否 | ✓ 1723ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1551ms | 否 | ✓ 1686ms | ✓ 1259ms | http |
| 20.78.118.91:8561 | ✓ 680ms | ✓ 1063ms | ✓ 795ms | ✓ 1023ms | ✓ 811ms | http |
| 20.210.39.153:8561 | ✓ 680ms | ✓ 1060ms | ✓ 797ms | ✓ 1027ms | 否 | http |
| 20.78.26.206:8561 | ✓ 685ms | ✓ 1101ms | ✓ 749ms | ✓ 1043ms | ✓ 780ms | http |
| 2.27.40.180:1080 | ✓ 446ms | ✓ 1386ms | ✓ 1306ms | ✓ 1727ms | ✓ 1454ms | http |
| 168.110.52.228:3128 | ✓ 1284ms | 否 | ✓ 1371ms | ✓ 1188ms | ✓ 1047ms | http |
| 150.107.140.238:3128 | 否 | 否 | ✓ 1255ms | ✓ 1853ms | ✓ 1119ms | http |
| 130.61.174.200:1080 | 否 | 否 | ✓ 1632ms | ✓ 1882ms | ✓ 1451ms | http |
| 45.129.141.143:3128 | ✓ 650ms | ✓ 1794ms | ✓ 1624ms | ✓ 1738ms | 否 | http |
| 138.124.99.216:8888 | ✓ 1899ms | ✓ 1583ms | ✓ 1466ms | 否 | ✓ 1432ms | http |
| 94.158.219.111:3128 | ✓ 1216ms | ✓ 1958ms | ✓ 1048ms | ✓ 1720ms | ✓ 1497ms | http |
| 121.230.8.89:1080 | ✓ 1102ms | ✓ 1489ms | ✓ 1128ms | ✓ 1480ms | 否 | http |
| 8.219.195.129:1080 | ✓ 1003ms | ✓ 1975ms | ✓ 1054ms | ✓ 1314ms | ✓ 1114ms | http |
| 91.233.223.147:3128 | ✓ 708ms | 否 | ✓ 1696ms | ✓ 1913ms | ✓ 1601ms | http |
| 152.32.132.190:7890 | ✓ 1340ms | 否 | 否 | ✓ 1962ms | ✓ 1579ms | http |
| 62.113.119.14:8080 | ✓ 522ms | ✓ 1598ms | ✓ 545ms | ✓ 1423ms | ✓ 1052ms | http |
| 115.231.181.40:8128 | ✓ 1484ms | ✓ 1328ms | 否 | ✓ 1464ms | 否 | http |
| 208.87.243.199:7878 | ✓ 1016ms | 否 | 否 | ✓ 1750ms | ✓ 1528ms | http |
| 209.126.10.139:3128 | ✓ 794ms | ✓ 1367ms | ✓ 1607ms | ✓ 1204ms | 否 | http |
| 104.248.243.244:3128 | ✓ 1350ms | ✓ 1779ms | ✓ 1523ms | 否 | 否 | http |
| 34.96.238.40:8080 | 否 | ✓ 1298ms | 否 | ✓ 1243ms | ✓ 1305ms | http |
| 37.187.109.70:10111 | ✓ 1067ms | ✓ 1496ms | ✓ 511ms | ✓ 1711ms | ✓ 1875ms | http |
| 43.132.188.134:443 | ✓ 1972ms | 否 | 否 | ✓ 1692ms | ✓ 1842ms | http |
| 146.19.56.212:40002 | ✓ 1906ms | ✓ 1451ms | ✓ 377ms | ✓ 1593ms | 否 | http |
| 94.131.118.129:1081 | ✓ 535ms | ✓ 1251ms | ✓ 1356ms | ✓ 1873ms | 否 | http |
| 94.131.118.129:1082 | ✓ 832ms | ✓ 1418ms | ✓ 738ms | 否 | 否 | http |
| 102.212.44.151:12354 | ✓ 1901ms | 否 | ✓ 1378ms | 否 | ✓ 1936ms | http |
| 131.186.27.113:10900 | ✓ 1074ms | 否 | 否 | ✓ 1284ms | ✓ 1204ms | http |
| 114.237.77.250:1080 | ✓ 1129ms | ✓ 1364ms | ✓ 1200ms | 否 | ✓ 1176ms | http |
| 59.3.228.41:65001 | ✓ 1377ms | ✓ 1665ms | ✓ 1613ms | 否 | 否 | http |
| 159.89.31.62:8080 | ✓ 388ms | ✓ 1701ms | ✓ 1595ms | ✓ 1833ms | 否 | http |
| 168.144.75.9:3128 | ✓ 1220ms | 否 | ✓ 1591ms | ✓ 1944ms | ✓ 1156ms | http |
| 152.42.177.32:8888 | ✓ 1192ms | 否 | ✓ 1735ms | ✓ 1564ms | ✓ 1619ms | http |
| 218.108.131.186:17890 | ✓ 1025ms | 否 | ✓ 1091ms | ✓ 1749ms | 否 | http |
| 18.170.25.193:9002 | ✓ 1614ms | 否 | ✓ 1104ms | 否 | ✓ 1647ms | http |
| 34.101.184.164:3128 | ✓ 1722ms | 否 | ✓ 1807ms | 否 | ✓ 1575ms | http |
| 168.222.254.136:8888 | ✓ 840ms | 否 | ✓ 1392ms | ✓ 1623ms | ✓ 1798ms | http |
| 8.209.238.110:47701 | ✓ 1758ms | 否 | ✓ 1961ms | ✓ 1110ms | ✓ 1044ms | http |
| 20.78.213.56:80 | ✓ 1569ms | 否 | 否 | ✓ 1826ms | ✓ 979ms | http |
| 42.101.8.101:8888 | 否 | 否 | ✓ 1358ms | ✓ 1701ms | ✓ 1337ms | http |
| 92.113.149.172:1080 | ✓ 1878ms | ✓ 1995ms | ✓ 879ms | 否 | 否 | http |
| 150.249.255.91:3128 | 否 | ✓ 1220ms | ✓ 687ms | 否 | ✓ 892ms | http |
| 43.252.238.138:8080 | 否 | 否 | ✓ 1953ms | ✓ 1777ms | ✓ 1431ms | http |
| 64.188.77.26:3128 | ✓ 1018ms | ✓ 1517ms | ✓ 592ms | ✓ 1659ms | 否 | http |
| 137.59.47.73:3128 | 否 | 否 | ✓ 1663ms | ✓ 1696ms | ✓ 1951ms | http |
| 61.52.131.172:8443 | ✓ 1148ms | ✓ 1374ms | ✓ 1188ms | ✓ 1462ms | ✓ 1165ms | http |
| 45.140.147.82:1082 | ✓ 961ms | ✓ 1365ms | 否 | ✓ 1597ms | ✓ 975ms | http |

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
