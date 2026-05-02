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

最后更新：2026-05-02 14:10:24 UTC（2026-05-02 22:10:24 UTC+8）

**代理总数：62**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 62 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 47.77.216.82:1080 | ✓ 955ms | 否 | ✓ 988ms | ✓ 1001ms | ✓ 684ms | http |
| 47.85.51.197:1080 | 否 | 否 | ✓ 444ms | ✓ 874ms | ✓ 1146ms | http |
| 218.108.131.186:17890 | ✓ 1040ms | ✓ 1266ms | ✓ 1031ms | ✓ 1398ms | ✓ 1070ms | http |
| 1.231.81.166:3128 | 否 | ✓ 1657ms | ✓ 1846ms | ✓ 1648ms | ✓ 1148ms | http |
| 113.160.132.26:8080 | 否 | ✓ 1540ms | ✓ 1480ms | ✓ 1802ms | ✓ 1166ms | http |
| 45.167.124.71:999 | ✓ 1420ms | ✓ 1772ms | ✓ 1401ms | ✓ 1887ms | ✓ 1625ms | http |
| 86.104.72.219:1081 | ✓ 246ms | ✓ 1600ms | ✓ 65ms | ✓ 1459ms | ✓ 903ms | http |
| 20.127.128.70:8080 | ✓ 675ms | 否 | ✓ 159ms | ✓ 1810ms | ✓ 1092ms | http |
| 107.150.41.226:18080 | 否 | 否 | ✓ 976ms | ✓ 1195ms | ✓ 977ms | http |
| 91.184.241.12:443 | ✓ 1597ms | 否 | ✓ 1695ms | 否 | ✓ 1571ms | http |
| 148.230.4.241:999 | ✓ 694ms | 否 | ✓ 652ms | ✓ 1765ms | ✓ 1491ms | http |
| 45.153.231.229:8080 | ✓ 1436ms | ✓ 1810ms | ✓ 1549ms | ✓ 1667ms | 否 | http |
| 20.164.75.153:8080 | ✓ 1642ms | 否 | ✓ 1880ms | 否 | ✓ 1798ms | http |
| 120.92.108.86:7890 | ✓ 1483ms | 否 | ✓ 1433ms | 否 | ✓ 1597ms | http |
| 149.51.42.10:3128 | ✓ 594ms | ✓ 1246ms | 否 | ✓ 1425ms | 否 | http |
| 86.104.72.219:1082 | ✓ 141ms | ✓ 1402ms | ✓ 48ms | 否 | ✓ 709ms | http |
| 62.113.119.14:8080 | ✓ 1050ms | 否 | ✓ 651ms | ✓ 1401ms | ✓ 1020ms | http |
| 133.242.16.174:3128 | ✓ 1206ms | 否 | 否 | ✓ 1348ms | ✓ 1931ms | http |
| 107.174.208.190:3128 | ✓ 1429ms | ✓ 1570ms | ✓ 860ms | ✓ 1215ms | ✓ 879ms | http |
| 206.206.126.177:2412 | ✓ 912ms | 否 | ✓ 1377ms | ✓ 1240ms | ✓ 970ms | http |
| 72.11.151.159:6005 | ✓ 1074ms | 否 | ✓ 1878ms | ✓ 1778ms | ✓ 1081ms | http |
| 103.157.200.126:3128 | ✓ 1867ms | 否 | ✓ 1191ms | 否 | ✓ 1231ms | http |
| 160.238.65.5:3128 | ✓ 1263ms | ✓ 1538ms | ✓ 1781ms | 否 | ✓ 1514ms | http |
| 135.125.97.184:38833 | ✓ 1378ms | 否 | ✓ 1471ms | 否 | ✓ 1872ms | http |
| 120.92.212.16:7890 | ✓ 1580ms | 否 | ✓ 1249ms | ✓ 1570ms | ✓ 1641ms | http |
| 117.236.124.166:3128 | ✓ 1207ms | 否 | ✓ 1015ms | 否 | ✓ 1457ms | http |
| 213.111.146.36:18080 | ✓ 1108ms | ✓ 1827ms | ✓ 1592ms | 否 | ✓ 1653ms | http |
| 34.101.184.164:3128 | ✓ 1910ms | 否 | ✓ 1573ms | ✓ 1542ms | ✓ 1260ms | http |
| 149.51.42.10:8080 | ✓ 638ms | ✓ 1238ms | 否 | ✓ 1196ms | 否 | http |
| 152.42.177.32:8888 | ✓ 1173ms | 否 | ✓ 1147ms | ✓ 1421ms | ✓ 1703ms | http |
| 103.35.190.69:1082 | ✓ 216ms | ✓ 956ms | ✓ 795ms | ✓ 1837ms | ✓ 734ms | http |
| 43.133.44.89:8888 | ✓ 1082ms | 否 | ✓ 1927ms | ✓ 1740ms | 否 | http |
| 72.11.150.178:6005 | ✓ 620ms | ✓ 1868ms | ✓ 823ms | ✓ 1344ms | ✓ 922ms | http |
| 150.107.140.238:3128 | ✓ 1945ms | 否 | ✓ 1145ms | ✓ 1330ms | 否 | http |
| 185.195.71.218:18080 | ✓ 521ms | ✓ 1944ms | ✓ 1698ms | 否 | ✓ 1704ms | http |
| 59.46.216.131:30001 | 否 | ✓ 1623ms | 否 | ✓ 1531ms | ✓ 1362ms | http |
| 223.84.151.86:30005 | 否 | ✓ 1867ms | ✓ 1559ms | ✓ 1764ms | ✓ 1783ms | http |
| 3.101.133.120:80 | ✓ 1855ms | 否 | ✓ 1380ms | ✓ 1687ms | ✓ 1229ms | http |
| 92.119.127.208:6005 | ✓ 471ms | 否 | ✓ 1262ms | ✓ 1684ms | 否 | http |
| 222.107.27.7:8056 | ✓ 1869ms | 否 | ✓ 1171ms | ✓ 1283ms | ✓ 969ms | http |
| 121.135.144.141:8052 | ✓ 1866ms | 否 | ✓ 1399ms | ✓ 1256ms | ✓ 1551ms | http |
| 212.58.132.5:8888 | ✓ 1066ms | 否 | ✓ 1868ms | ✓ 1508ms | ✓ 1210ms | http |
| 77.232.142.164:3128 | ✓ 1437ms | 否 | 否 | ✓ 1886ms | ✓ 1550ms | http |
| 168.110.52.228:3128 | ✓ 1843ms | 否 | 否 | ✓ 1249ms | ✓ 1697ms | http |
| 130.61.174.200:1080 | 否 | 否 | ✓ 586ms | ✓ 1414ms | ✓ 1993ms | http |
| 223.16.170.103:3128 | 否 | 否 | ✓ 1372ms | ✓ 1756ms | ✓ 1405ms | http |
| 107.174.80.186:3128 | 否 | ✓ 1289ms | 否 | ✓ 1105ms | ✓ 998ms | http |
| 104.248.243.244:3128 | ✓ 990ms | 否 | ✓ 1168ms | ✓ 1745ms | ✓ 1233ms | http |
| 34.96.238.40:8080 | ✓ 1461ms | 否 | ✓ 1033ms | 否 | ✓ 1234ms | http |
| 62.60.231.71:56608 | 否 | 否 | ✓ 658ms | ✓ 1886ms | ✓ 1295ms | http |
| 220.197.44.36:3128 | ✓ 1565ms | ✓ 1993ms | ✓ 1830ms | ✓ 1908ms | ✓ 1644ms | http |
| 103.35.190.69:1081 | ✓ 215ms | ✓ 1591ms | ✓ 66ms | ✓ 945ms | ✓ 684ms | http |
| 89.208.106.138:10808 | ✓ 991ms | ✓ 1855ms | ✓ 406ms | 否 | ✓ 902ms | http |
| 101.32.243.189:80 | ✓ 1948ms | 否 | ✓ 1436ms | ✓ 1745ms | ✓ 1483ms | http |
| 61.52.131.172:8443 | 否 | ✓ 1971ms | ✓ 1136ms | ✓ 1410ms | ✓ 1105ms | http |
| 45.129.141.143:3128 | ✓ 1599ms | 否 | ✓ 1747ms | ✓ 1992ms | ✓ 1680ms | http |
| 152.70.91.193:40000 | ✓ 1788ms | 否 | 否 | ✓ 1468ms | ✓ 1627ms | http |
| 20.210.39.153:8561 | 否 | 否 | ✓ 943ms | ✓ 1027ms | ✓ 819ms | http |
| 20.78.26.206:8561 | 否 | 否 | ✓ 945ms | ✓ 1054ms | ✓ 780ms | http |
| 94.131.118.129:1081 | ✓ 1202ms | ✓ 1503ms | ✓ 1001ms | ✓ 1559ms | ✓ 1489ms | http |
| 223.16.170.103:80 | ✓ 1342ms | 否 | ✓ 1343ms | ✓ 1338ms | ✓ 1336ms | http |
| 103.166.182.144:3128 | ✓ 1976ms | 否 | ✓ 1105ms | ✓ 1312ms | ✓ 1063ms | http |

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
