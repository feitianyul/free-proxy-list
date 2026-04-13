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

最后更新：2026-04-13 12:09:01 UTC（2026-04-13 20:09:01 UTC+8）

**代理总数：49**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 49 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 49 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 43.156.132.113:3128 | ✓ 1644ms | ✓ 1831ms | ✓ 930ms | ✓ 1267ms | ✓ 985ms | http |
| 167.103.115.102:8800 | ✓ 1649ms | 否 | 否 | ✓ 1556ms | ✓ 1527ms | http |
| 167.103.34.108:8800 | ✓ 1682ms | 否 | ✓ 1614ms | ✓ 1910ms | ✓ 1772ms | http |
| 147.161.210.140:8800 | ✓ 1678ms | 否 | ✓ 963ms | ✓ 1441ms | ✓ 1287ms | http |
| 160.250.134.143:3128 | ✓ 1168ms | 否 | 否 | ✓ 1440ms | ✓ 1145ms | http |
| 167.103.144.127:8800 | ✓ 1659ms | 否 | 否 | ✓ 1983ms | ✓ 1845ms | http |
| 46.39.105.157:8080 | ✓ 1012ms | 否 | ✓ 1613ms | 否 | ✓ 1383ms | http |
| 45.167.124.52:8080 | 否 | 否 | ✓ 514ms | ✓ 1593ms | ✓ 1206ms | http |
| 45.167.125.21:999 | 否 | 否 | ✓ 1305ms | ✓ 1700ms | ✓ 1353ms | http |
| 113.160.132.26:8080 | ✓ 1887ms | 否 | ✓ 1298ms | ✓ 1664ms | 否 | http |
| 167.103.31.122:8800 | ✓ 1282ms | 否 | ✓ 1407ms | ✓ 1607ms | ✓ 1454ms | http |
| 35.225.22.61:80 | ✓ 389ms | ✓ 1266ms | 否 | ✓ 1543ms | 否 | http |
| 45.149.92.147:5001 | 否 | 否 | ✓ 1105ms | ✓ 1067ms | ✓ 926ms | http |
| 139.159.99.242:8080 | 否 | 否 | ✓ 1083ms | ✓ 1272ms | ✓ 1248ms | http |
| 103.157.200.126:3128 | 否 | 否 | ✓ 1442ms | ✓ 1605ms | ✓ 1272ms | http |
| 114.237.77.202:1080 | 否 | ✓ 1461ms | ✓ 1603ms | ✓ 1403ms | ✓ 1247ms | http |
| 121.43.196.213:8222 | 否 | ✓ 1314ms | ✓ 1114ms | ✓ 1315ms | ✓ 1051ms | http |
| 121.43.196.210:8222 | ✓ 1046ms | ✓ 1308ms | ✓ 1000ms | 否 | 否 | http |
| 159.223.225.118:8888 | ✓ 858ms | 否 | 否 | ✓ 1561ms | ✓ 1278ms | http |
| 129.213.162.27:17777 | 否 | ✓ 1828ms | ✓ 1770ms | ✓ 1737ms | 否 | http |
| 137.59.47.73:3128 | ✓ 1793ms | 否 | ✓ 1680ms | ✓ 1363ms | 否 | http |
| 2.27.32.81:3128 | ✓ 1260ms | ✓ 1865ms | ✓ 919ms | 否 | 否 | http |
| 120.92.108.86:7890 | ✓ 1528ms | 否 | ✓ 1495ms | 否 | ✓ 1532ms | http |
| 79.132.136.58:3128 | ✓ 446ms | 否 | ✓ 397ms | ✓ 1194ms | ✓ 944ms | http |
| 8.209.238.110:47701 | 否 | 否 | ✓ 826ms | ✓ 1124ms | ✓ 878ms | http |
| 101.32.244.83:8080 | ✓ 1148ms | 否 | ✓ 1129ms | ✓ 1538ms | ✓ 1460ms | http |
| 114.55.226.123:10086 | ✓ 1310ms | ✓ 1642ms | ✓ 1248ms | ✓ 1557ms | ✓ 1232ms | http |
| 45.186.6.104:3128 | ✓ 1514ms | ✓ 1895ms | ✓ 1858ms | 否 | 否 | http |
| 129.212.224.122:3128 | 否 | 否 | ✓ 1010ms | ✓ 1263ms | ✓ 1004ms | http |
| 8.219.97.248:80 | ✓ 1285ms | 否 | ✓ 1626ms | ✓ 1491ms | 否 | http |
| 103.82.23.118:5171 | ✓ 1664ms | 否 | ✓ 1582ms | 否 | ✓ 1752ms | http |
| 62.113.119.14:8080 | 否 | ✓ 1708ms | ✓ 741ms | ✓ 1919ms | 否 | http |
| 218.108.131.186:17890 | ✓ 998ms | ✓ 1270ms | ✓ 1083ms | ✓ 1326ms | ✓ 1056ms | http |
| 147.161.239.240:8800 | ✓ 1068ms | 否 | ✓ 1000ms | ✓ 1573ms | ✓ 1368ms | http |
| 5.104.87.17:8051 | 否 | 否 | ✓ 1580ms | ✓ 1749ms | ✓ 1365ms | http |
| 212.58.132.5:8888 | 否 | 否 | ✓ 1865ms | ✓ 1576ms | ✓ 1658ms | http |
| 110.42.37.202:20005 | ✓ 1969ms | ✓ 1787ms | ✓ 1320ms | ✓ 1727ms | ✓ 1478ms | http |
| 217.76.245.80:999 | 否 | ✓ 1372ms | ✓ 994ms | ✓ 1897ms | ✓ 1122ms | http |
| 34.50.41.219:3128 | 否 | 否 | ✓ 1498ms | ✓ 1548ms | ✓ 1903ms | http |
| 210.223.44.230:3128 | ✓ 1987ms | 否 | 否 | ✓ 1718ms | ✓ 1641ms | http |
| 46.30.46.133:3128 | ✓ 669ms | 否 | ✓ 1689ms | ✓ 1293ms | ✓ 966ms | http |
| 195.26.224.49:3128 | 否 | 否 | ✓ 1147ms | ✓ 1347ms | ✓ 1383ms | http |
| 8.219.195.129:1080 | 否 | 否 | ✓ 910ms | ✓ 1276ms | ✓ 981ms | http |
| 180.103.19.53:1080 | 否 | 否 | ✓ 1458ms | ✓ 1862ms | ✓ 1556ms | http |
| 101.43.127.100:8877 | 否 | ✓ 1802ms | ✓ 1999ms | 否 | ✓ 1898ms | http |
| 181.78.44.63:999 | ✓ 1535ms | 否 | ✓ 1038ms | ✓ 1329ms | ✓ 1073ms | http |
| 150.249.255.91:3128 | ✓ 1543ms | 否 | ✓ 914ms | ✓ 1420ms | 否 | http |
| 103.147.77.254:7777 | ✓ 1981ms | 否 | 否 | ✓ 1755ms | ✓ 1715ms | http |
| 138.197.68.35:4857 | ✓ 1796ms | ✓ 1166ms | ✓ 107ms | ✓ 914ms | ✓ 902ms | http |

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
